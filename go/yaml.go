package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/candiddev/shared/go/errs"
	"github.com/candiddev/shared/go/logger"
	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"
)

var errValidationFailure = errors.New("validation failure")

type translations struct {
	DefaultCode  string                       `json:"defaultCode"`
	ISO639Codes  map[string]string            `json:"iso639Codes"`
	Outputs      []output                     `json:"outputs"`
	Translations map[string]map[string]string `json:"translations"`
}

type output struct {
	Format   string `json:"format"`
	Package  string `json:"package"`
	Path     string `json:"path"`
	realPath string
}

func parseTranslations(ctx context.Context, c *config) (*translations, error) {
	t := translations{}

	content, err := os.ReadFile(c.Input)
	if err != nil {
		return nil, logger.Error(ctx, errs.ErrReceiver.Wrap(err))
	}

	if err := yaml.UnmarshalStrict(content, &t); err != nil {
		return nil, logger.Error(ctx, errs.ErrReceiver.Wrap(err))
	}

	if len(t.Outputs) == 0 {
		return nil, logger.Error(ctx, errs.ErrReceiver.Wrap(errors.New("no outputs defined")))
	}

	for i := range t.Outputs {
		if _, ok := formats[t.Outputs[i].Format]; !ok {
			return nil, logger.Error(ctx, errs.ErrReceiver.Wrap(fmt.Errorf("unknown format: %s", t.Outputs[i].Format)))
		}

		if t.Outputs[i].Package == "" {
			t.Outputs[i].Package = "yaml8n"
		}

		if t.Outputs[i].Path == "" {
			return nil, logger.Error(ctx, errs.ErrReceiver.Wrap(errors.New("output path must be specified")))
		}

		if strings.HasPrefix(t.Outputs[i].Path, "/") {
			t.Outputs[i].realPath = t.Outputs[i].Path
		} else {
			t.Outputs[i].realPath = filepath.Join(filepath.Dir(c.Input), t.Outputs[i].Path)
		}
	}

	logger.Error(ctx, nil) //nolint:errcheck

	return &t, t.validate(ctx, c, string(content))
}

func (t *translations) generate(ctx context.Context) error {
	for i := range t.Outputs {
		f, ok := formats[t.Outputs[i].Format]
		if !ok {
			return logger.Error(ctx, errs.ErrReceiver.Wrap(fmt.Errorf("unknown language: %s", t.Outputs[i].Format)))
		}

		if err := os.MkdirAll(t.Outputs[i].realPath, 0755); err != nil { //nolint:gosec
			return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
		}

		tmp, err := template.New("code").Funcs(template.FuncMap{
			"package": func() string {
				return t.Outputs[i].Package
			},
		}).Parse(f.Template)
		if err != nil {
			return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
		}

		if f.PerLanguage {
			for j := range t.ISO639Codes {
				t.DefaultCode = j

				if err := writeTemplate(ctx, t, tmp, filepath.Join(t.Outputs[i].realPath, fmt.Sprintf("%s.%s", j, f.Extension))); err != nil {
					return logger.Error(ctx, err)
				}
			}
		} else {
			if err := writeTemplate(ctx, t, tmp, filepath.Join(t.Outputs[i].realPath, fmt.Sprintf("%s.%s", t.Outputs[i].Package, f.Extension))); err != nil {
				return logger.Error(ctx, err)
			}
		}
	}

	return logger.Error(ctx, nil)
}

func writeTemplate(ctx context.Context, t *translations, tmp *template.Template, path string) errs.Err {
	var out bytes.Buffer

	if err := tmp.Execute(&out, t); err != nil {
		return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
	}

	if err := os.WriteFile(path, out.Bytes(), 0644); err != nil { //nolint: gosec
		return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
	}

	return nil
}

func (t *translations) validate(ctx context.Context, c *config, raw string) errs.Err { //nolint: gocognit,gocyclo
	var err error

	o := struct {
		Errs     []string
		Warnings []string
	}{
		Errs:     []string{},
		Warnings: []string{},
	}

	if t.DefaultCode == "" {
		o.Errs = append(o.Errs, "defaultCode must be specified")
	}

	if len(t.ISO639Codes) == 0 {
		o.Errs = append(o.Errs, "iso639Codes must be specified")
	}

	codes := []string{}
	checkMatch := false
	defaultMatch := false

	for k := range t.ISO639Codes {
		if k == t.DefaultCode {
			defaultMatch = true
		}

		if k == c.CheckCode {
			checkMatch = true
		}

		codes = append(codes, k)
	}

	if c.CheckCode != "" && !checkMatch {
		o.Errs = append(o.Errs, c.CheckCode+" must be specified in iso639Codes")
	}

	if !defaultMatch && t.DefaultCode != "" {
		o.Errs = append(o.Errs, t.DefaultCode+" must be specified in iso639Codes")
	}

	for k, v := range t.Translations {
		if _, ok := v["context"]; !ok && c.CheckCode == "" {
			o.Errs = append(o.Errs, k+" must have a context specified")
		}

		cs := append([]string{}, codes...)

		if c.CheckCode != "" {
			cs = []string{
				c.CheckCode,
			}
		}

		for code := range v {
			match := false

			if code == "context" {
				continue
			}

			for i := range cs {
				if cs[i] == code {
					cs = append(cs[:i], cs[i+1:]...)
					match = true

					break
				}
			}

			if !match && c.CheckCode == "" {
				o.Errs = append(o.Errs, fmt.Sprintf(`%s has a code not present in iso639Codes: %s`, k, code))
			}
		}

		for i := range cs {
			msg := fmt.Sprintf(`%s is missing a translation: %s`, k, cs[i])

			if cs[i] == t.DefaultCode {
				o.Errs = append(o.Errs, msg)
			} else {
				o.Warnings = append(o.Warnings, msg)
			}
		}
	}

	if raw != "" {
		y, err := yaml.Marshal(t)
		if err == nil {
			if d := cmp.Diff(string(y), raw); d != "" {
				o.Warnings = append(o.Warnings, "YAML not formatted correctly (got +, want -):\n"+strings.ReplaceAll(d, "\u00a0", ""))
			}
		} else {
			o.Errs = append(o.Errs, fmt.Sprintf("Error marshaling YAML: %s", err))
		}
	}

	if len(o.Errs) > 0 || len(o.Warnings) > 0 {
		var out bytes.Buffer

		if err := template.Must(template.New("report").Parse(`
{{- if .Errs -}}
{{ len .Errs }} Error{{ if ne (len .Errs) 1 }}s{{ end }} Found:
{{- range .Errs }}
- {{ . }}
{{- end }}
{{ end -}}
{{- if .Warnings -}}
{{ len .Warnings }} Warning{{ if ne (len .Warnings) 1 }}s{{ end }} Found:
{{- range .Warnings }}
- {{ . }}
{{- end }}
{{ end -}}
`)).Execute(&out, o); err != nil {
			return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
		}

		if len(o.Errs) > 0 || (len(o.Warnings) > 0 && c.FailWarn) {
			err = errValidationFailure
		}

		logger.Info(ctx, "\n", out.String())

		if err != nil {
			return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
		}
	}

	return nil
}
