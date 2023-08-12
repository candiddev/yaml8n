package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"sigs.k8s.io/yaml"
)

func (t *translations) runTranslate(c *config) error { //nolint: gocognit
	characters := 0
	codes := []string{}

	if c.CheckCode == "" {
		for k := range t.ISO639Codes {
			codes = append(codes, k)
		}
	} else {
		codes = []string{
			c.CheckCode,
		}
	}

	for _, v := range t.Translations {
		c := append([]string{}, codes...)

		for code := range v {
			if code == "context" { //nolint: goconst
				continue
			}

			for i := range c {
				if c[i] == code {
					c = append(c[:i], c[i+1:]...)

					break
				}
			}
		}

		for range c {
			characters += len(v[t.DefaultCode])
		}
	}

	if characters != 0 {
		fmt.Printf("Need to translate %d characters (type OK to continue): ", characters) //nolint: forbidigo

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if text != "OK\n" {
			fmt.Println("Cancelled") //nolint: forbidigo

			return nil
		}

		ctx := context.Background()

		l := len(t.Translations)
		n := 0

		client, err := translate.NewClient(ctx)
		if err != nil {
			fmt.Println(err) //nolint: forbidigo

			return err
		}
		defer client.Close()

		for _, v := range t.Translations {
			c := append([]string{}, codes...)
			n++

			for code := range v {
				if code == "context" {
					continue
				}

				for i := range c {
					if c[i] == code {
						c = append(c[:i], c[i+1:]...)

						break
					}
				}
			}

			for i := range c {
				fmt.Printf("Translating %d/%d - %s...\n", n, l, c[i]) //nolint: forbidigo

				l, err := language.Parse(c[i])
				if err != nil {
					fmt.Println(err) //nolint: forbidigo

					continue
				}

				resp, err := client.Translate(ctx, []string{v[t.DefaultCode]}, l, nil)
				if err != nil {
					fmt.Println(err) //nolint: forbidigo

					continue
				}

				if len(resp) != 1 {
					fmt.Println("Translation request was invalid or empty") //nolint: forbidigo

					continue
				}

				v[c[i]] = strings.ReplaceAll(resp[0].Text, "\u00a0", "")
			}
		}
	}

	y, err := yaml.Marshal(t)
	if err != nil {
		fmt.Println(err) //nolint: forbidigo

		return err
	}

	return os.WriteFile(c.Input, y, 0644) //nolint: gosec
}
