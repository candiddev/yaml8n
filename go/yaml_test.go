package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
)

func TestParseTranslations(t *testing.T) {
	tests := map[string]struct {
		err   string
		input *config
		want  *translations
	}{
		"no file": {
			err: "open ./notafile: no such file or directory",
			input: &config{
				Input: "./notafile",
			},
		},
		"bad yaml": {
			err: "error converting YAML to JSON: yaml: did not find expected key",
			input: &config{
				Input: "./testdata/bad.yaml",
			},
		},
		"no outputs": {
			err: "no outputs",
			input: &config{
				Input: "./testdata/badformat.yaml",
			},
		},
		"no path": {
			err: "output path must be",
			input: &config{
				Input: "./testdata/nopath.yaml",
			},
		},
		"invalid format": {
			err: "validation failure",
			input: &config{
				Input: "./testdata/invalid.yaml",
			},
			want: &translations{
				ISO639Codes: map[string]string{
					"en": "English",
				},
				Outputs: []output{
					{
						Format:  "go",
						Package: "yaml8n",
						Path:    "../",
					},
				},
				Translations: map[string]map[string]string{
					"hello": nil,
				},
			},
		},
		"valid format": {
			input: &config{
				Input: "./testdata/valid.yaml",
			},
			want: &translations{
				DefaultCode: "en",
				Outputs: []output{
					{
						Format:  "go",
						Package: "yaml8n",
						Path:    "../",
					},
				},
				ISO639Codes: map[string]string{
					"en": "English",
				},
				Translations: map[string]map[string]string{
					"hello": {
						"context": "Context",
						"en":      "Hello!",
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseTranslations(ctx, tc.input)

			if tc.err == "" {
				assert.HasErr(t, err, nil)
			} else {
				assert.Equal(t, strings.Contains(err.Error(), tc.err), true)
			}

			assert.EqualJSON(t, got, tc.want)
		})
	}
}

func testTranslationsGenerate(t *testing.T, format string, outputs map[string]string) {
	t.Helper()
	logger.UseTestLogger(t)

	f := "testdata/translations"

	tra := translations{
		DefaultCode: "en",
		ISO639Codes: map[string]string{
			"de": "Deutsch",
			"en": "English",
		},
		Outputs: []output{
			{
				Format:   format,
				Path:     f,
				Package:  "yaml8n",
				realPath: f,
			},
		},
		Translations: map[string]map[string]string{
			"HelloWorld": {
				"de": "Hallo Welt",
				"en": "Hello World",
			},
			"TranslateMe": {
				"en": "Translate Me",
			},
		},
	}

	tra.generate(ctx)

	for path, content := range outputs {
		s, err := os.ReadFile(filepath.Join(f, path))

		assert.Equal(t, err, nil)
		assert.Equal(t, string(s), content)

		assert.Equal(t, err, nil)
	}

	os.RemoveAll(f)
}

func TestTranslationsValidate(t *testing.T) {
	tests := map[string]struct {
		checkCode string
		err       error
		failWarn  bool
		raw       string
		tra       translations
		want      string
	}{
		"no default code": {
			err: errValidationFailure,
			tra: translations{
				ISO639Codes: map[string]string{
					"en": "English",
				},
			},
			want: `1 Error Found:
- defaultCode must be specified
`,
		},
		"no iso codes": {
			err: errValidationFailure,
			tra: translations{
				DefaultCode: "en",
			},
			want: `2 Errors Found:
- iso639Codes must be specified
- en must be specified in iso639Codes
`,
		},
		"missing checkCode": {
			checkCode: "de",
			err:       errValidationFailure,
			raw: `defaultcode: en
iso639codes:
  en: English
translations: null
`,
			tra: translations{
				DefaultCode: "en",
				ISO639Codes: map[string]string{
					"en": "English",
				},
			},
			want: `1 Error Found:
- de must be specified in iso639Codes
`,
		},
		"missing translations": {
			err: errValidationFailure,
			tra: translations{
				DefaultCode: "en",
				ISO639Codes: map[string]string{
					"de": "Deutsch",
					"en": "English",
				},
				Translations: map[string]map[string]string{
					"HelloWorld": {},
				},
			},
			want: `2 Errors Found:
- HelloWorld must have a context specified
- HelloWorld is missing a translation: en
1 Warning Found:
- HelloWorld is missing a translation: de
`,
		},
		"extra code": {
			err: errValidationFailure,
			tra: translations{
				DefaultCode: "en",
				ISO639Codes: map[string]string{
					"en": "English",
				},
				Translations: map[string]map[string]string{
					"HelloWorld": {
						"context": "context",
						"en":      "Hello World",
						"de":      "Hello World",
					},
				},
			},
			want: `1 Error Found:
- HelloWorld has a code not present in iso639Codes: de
`,
		},
		"yaml fail": {
			tra: translations{
				DefaultCode: "en",
				ISO639Codes: map[string]string{
					"en": "English",
				},
			},
			raw: `defaultCODE: "en"
iso639codes:
  en: English
translations: null
`,
			want: `1 Warning Found:
- YAML not formatted correctly (got +, want -):
`,
		},
		"warnings": {
			tra: translations{
				DefaultCode: "en",
				ISO639Codes: map[string]string{
					"de": "Deutsch",
					"en": "English",
				},
				Translations: map[string]map[string]string{
					"HelloWorld": {
						"context": "context",
						"en":      "Hello World",
					},
				},
			},
			want: `1 Warning Found:
- HelloWorld is missing a translation: de
`,
		},
		"warnings fail": {
			err:      errValidationFailure,
			failWarn: true,
			tra: translations{
				DefaultCode: "en",
				ISO639Codes: map[string]string{
					"de": "Deutsch",
					"en": "English",
				},
				Translations: map[string]map[string]string{
					"HelloWorld": {
						"context": "context",
						"en":      "Hello World",
					},
				},
			},
			want: `1 Warning Found:
- HelloWorld is missing a translation: de
`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			logger.SetStd()
			err := tc.tra.validate(ctx, &config{
				CheckCode: tc.checkCode,
				FailWarn:  tc.failWarn,
			}, tc.raw)

			assert.HasErr(t, err, tc.err)
			assert.Equal(t, strings.Contains(logger.ReadStd(), tc.want), true)
		})
	}
}
