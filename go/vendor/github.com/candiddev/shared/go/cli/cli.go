// Package cli contains functions for building CLIs.
package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/candiddev/shared/go/config"
	"github.com/candiddev/shared/go/errs"
	"github.com/candiddev/shared/go/logger"
	"golang.org/x/term"
)

// BuildDate is the application build date in YYYY-MM-DD, set with candid/lib/cli.Builddate build time variable.
var BuildDate string //nolint:gochecknoglobals

// BuildVersion is the application version, set with candid/lib/cli.BuildVersion build time variable.
var BuildVersion string //nolint:gochecknoglobals

// Config manages the CLI configuration.
type Config struct {
	ConfigPath    string        `json:"configPath"`
	LogFormat     logger.Format `json:"logFormat"`
	LogLevel      logger.Level  `json:"logLevel"`
	NoColor       bool          `json:"noColor"`
	runMock       *runMock
	runMockEnable bool
}

type runMock struct {
	inputs  []RunMockInput
	errs    []error
	mutex   *sync.Mutex
	outputs []string
}

// Command is a positional command to run.
type Command[T AppConfig[any]] struct {
	/* Optional Positional arguments after command */
	ArgumentsOptional []string

	/* Positional arguments required after command */
	ArgumentsRequired []string

	/* Override the command name in usage */
	Name string

	/* Function to run when calling the command */
	Run func(ctx context.Context, args []string, config T) errs.Err

	/* Usage information, omitting this hides the command */
	Usage string
}

var ErrUnknownCommand = errs.ErrSenderNotFound.Wrap(errors.New("unknown command"))

// ConfigArgs is a list of config arguments.
type ConfigArgs []string

func (i *ConfigArgs) Set(value string) error {
	*i = append(*i, value)

	return nil
}

func (i *ConfigArgs) String() string {
	return strings.Join(*i, "")
}

// App is a CLI application.
type App[T AppConfig[any]] struct {
	Commands         map[string]Command[T]
	Config           T
	Description      string
	HideConfigFields []string
	Name             string
	NoParse          bool
}

func wrapLines(l int, lines string, indent string) string {
	s := strings.Fields(strings.TrimSpace(lines))
	if len(s) == 0 {
		return lines
	}

	o := s[0]
	n := l - len(o)

	for _, w := range s[1:] {
		if len(w)+1 > n {
			o += "\n" + indent + w
			n = l - len(w)
		} else {
			o += " " + w
			n -= 1 + len(w)
		}
	}

	return o
}

// AppConfig is a configuration that can be used with CLI.
type AppConfig[T any] interface {
	CLIConfig() *Config
	Parse(ctx context.Context, configArgs []string) errs.Err
}

// Run is the main entrypoint into a CLI app.
func (a App[T]) Run() errs.Err { //nolint:gocognit
	ctx := context.Background()

	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.Usage = func() {}
	f.SetOutput(logger.Stdout)

	usage := func(arg string) {
		if arg == "" {
			//nolint:forbidigo
			fmt.Fprintf(logger.Stdout, `Usage: %s [flags] [command]

%s

Commands:
`, a.Name, a.Description)
		} else {
			//nolint:forbidigo
			fmt.Fprintln(logger.Stdout)
		}

		c := []string{}

		for i := range a.Commands {
			if a.Commands[i].Usage != "" {
				c = append(c, i)
			}
		}

		sort.Strings(c)

		w, _, _ := term.GetSize(0)

		if w == 0 || w > 70 {
			w = 70
		}

		for i := range c {
			name := c[i]
			if (a.Commands[c[i]]).Name != "" {
				name = a.Commands[c[i]].Name
			}

			if arg != "" && arg != name {
				continue
			}

			for _, arg := range a.Commands[c[i]].ArgumentsRequired {
				name += fmt.Sprintf(" [%s]", arg)
			}

			for _, arg := range a.Commands[c[i]].ArgumentsOptional {
				name += fmt.Sprintf(" [%s]", arg)
			}

			fmt.Fprintf(logger.Stdout, "  %s\n    	%s\n\n", wrapLines(w, name, ""), wrapLines(w, a.Commands[c[i]].Usage, "     	")) //nolint:forbidigo
		}

		//nolint: forbidigo
		fmt.Fprintf(logger.Stdout, "Flags:\n")
		f.PrintDefaults()
	}

	a.Commands["jq"] = Command[T]{
		ArgumentsOptional: []string{
			"-r, render raw values",
			"query string",
		},
		Run:   jq[T],
		Usage: "Query JSON from stdin using jq.  Supports standard JQ queries.",
	}

	c := ConfigArgs{}

	if !a.NoParse {
		a.Config.CLIConfig().ConfigPath = strings.ToLower(a.Name) + ".jsonnet"

		f.StringVar(&a.Config.CLIConfig().ConfigPath, "c", a.Config.CLIConfig().ConfigPath, "Path to JSON/Jsonnet configuration files separated by a comma")

		a.Commands["show-config"] = Command[T]{
			Run: func(ctx context.Context, args []string, config T) errs.Err {
				return printConfig(ctx, a)
			},
			Usage: "Print the current configuration",
		}

		f.Var(&c, "x", "Set config key=value (can be provided multiple times)")
	}

	a.Commands["version"] = Command[T]{
		Run: func(ctx context.Context, args []string, config T) errs.Err {
			fmt.Fprintf(logger.Stdout, "Build Version: %s\n", BuildVersion) //nolint: forbidigo
			fmt.Fprintf(logger.Stdout, "Build Date: %s\n", BuildDate)       //nolint: forbidigo

			return nil
		},
		Usage: "Print version information",
	}

	f.StringVar((*string)(&a.Config.CLIConfig().LogFormat), "f", string(a.Config.CLIConfig().LogFormat), "Set log format (human, kv, raw, default: human)")
	f.StringVar((*string)(&a.Config.CLIConfig().LogLevel), "l", string(a.Config.CLIConfig().LogLevel), "Set minimum log level (none, debug, info, error, default: info)")
	f.BoolVar(&a.Config.CLIConfig().NoColor, "n", a.Config.CLIConfig().NoColor, "Disable colored logging")

	if err := f.Parse(os.Args[1:]); err != nil {
		usage("")

		return ErrUnknownCommand
	}

	// Parse CLI environment early for logging options.
	if err := config.ParseValues(ctx, a.Config, strings.ToUpper(a.Name)+"_cli_", os.Environ()); err != nil {
		return logger.Error(ctx, errs.ErrReceiver.Wrap(config.ErrUpdateEnv, err))
	}

	ctx = logger.SetFormat(ctx, a.Config.CLIConfig().LogFormat)
	ctx = logger.SetLevel(ctx, a.Config.CLIConfig().LogLevel)
	ctx = logger.SetNoColor(ctx, a.Config.CLIConfig().NoColor)

	if !a.NoParse {
		// Resolve the real config path early by walking parent directories.  If the real config path exists (isn't "") and is different than the current one, update the path value.
		if p := config.FindPathAscending(ctx, a.Config.CLIConfig().ConfigPath); p != "" && p != a.Config.CLIConfig().ConfigPath {
			a.Config.CLIConfig().ConfigPath = p
		}

		if err := a.Config.Parse(ctx, c); err != nil {
			return err
		}
	}

	// Refresh ctx for new config values.
	ctx = logger.SetFormat(ctx, a.Config.CLIConfig().LogFormat)
	ctx = logger.SetLevel(ctx, a.Config.CLIConfig().LogLevel)
	ctx = logger.SetNoColor(ctx, a.Config.CLIConfig().NoColor)

	args := f.Args()
	if len(args) < 1 {
		usage("")

		return ErrUnknownCommand
	}

	for k, v := range a.Commands {
		if k == args[0] || strings.Split(v.Name, " ")[0] == args[0] {
			if len(v.ArgumentsRequired) != 0 && (len(args)-1) < len(v.ArgumentsRequired) {
				logger.Error(ctx, errs.ErrReceiver.Wrap(errors.New("missing arguments: ["+strings.Join(v.ArgumentsRequired[0+len(args)-1:], "] [")+"]\n"))) //nolint:errcheck

				usage(args[0])

				return ErrUnknownCommand
			}

			return v.Run(ctx, args, a.Config)
		}
	}

	usage("")

	return ErrUnknownCommand
}
