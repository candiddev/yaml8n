// yaml8n is a CLI tool for generating translations.
package main

import (
	"flag"
	"os"

	"github.com/candiddev/shared/go/cli"
)

func main() {
	c := &config{}

	flag.StringVar(&c.CheckCode, "c", c.CheckCode, "Check a specific language code")
	flag.BoolVar(&c.FailWarn, "w", c.FailWarn, "Fail on translation warnings")

	if err := (cli.App[*config]{
		Commands: map[string]cli.Command[*config]{
			"generate": {
				ArgumentsRequired: []string{
					"path",
				},
				Run:   run,
				Usage: "Generate code for programming languages",
			},
			"translate": {
				ArgumentsRequired: []string{
					"path",
				},
				Run:   run,
				Usage: "Create new translations using Google Cloud Translation API",
			},
			"validate": {
				ArgumentsRequired: []string{
					"path",
				},
				Run:   run,
				Usage: "Validate the translations",
			},
			"watch": {
				ArgumentsRequired: []string{
					"path",
				},
				Run:   run,
				Usage: "Watch the input file and regenerate code on changes",
			},
		},
		Config:      c,
		Description: "YAML8n makes translating your app easy and type safe.",
		Name:        "YAML8n",
		NoParse:     true,
	}).Run(); err != nil {
		os.Exit(1)
	}
}
