package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/candiddev/shared/go/errs"
	"github.com/candiddev/shared/go/logger"
	"github.com/fsnotify/fsnotify"
)

func run(ctx context.Context, args []string, c *config) errs.Err { //nolint: gocognit
	c.Input = args[1]

	if args[0] == "watch" {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err) //nolint:revive
		}
		defer watcher.Close()

		go func() {
			var lastEvent *fsnotify.Event

			timer := time.NewTimer(time.Millisecond)
			<-timer.C // timer should be expired at first

			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}

					if event.Has(fsnotify.Write) {
						lastEvent = &event

						timer.Reset(time.Millisecond * 100)
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}

					fmt.Println(err) //nolint: forbidigo
				case <-timer.C:
					if lastEvent != nil {
						fmt.Println(lastEvent) //nolint: forbidigo

						t, err := parseTranslations(ctx, c)
						if err != nil {
							continue
						}

						if err := t.generate(ctx); err != nil {
							continue
						}

						fmt.Println() //nolint: forbidigo
					}

					lastEvent = nil
				}
			}
		}()

		err = watcher.Add(c.Input)
		if err != nil {
			return logger.Log(ctx, errs.NewCLIErr(fmt.Errorf("error opening %s: %w", c.Input, err)))
		}

		// Block main goroutine forever.
		<-make(chan struct{})
	} else {
		t, err := parseTranslations(ctx, c)
		if err != nil {
			os.Exit(1) //nolint:gocritic,revive
		}

		if args[0] == "validate" {
			os.Exit(0) //nolint:revive
		} else if args[0] == "translate" {
			if err := t.runTranslate(c); err != nil {
				os.Exit(1) //nolint:revive
			}

			os.Exit(0) //nolint:revive
		}

		if err := t.generate(ctx); err != nil {
			os.Exit(1) //nolint:revive
		}
	}

	return nil
}
