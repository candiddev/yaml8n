package main

import (
	"context"

	"github.com/candiddev/shared/go/cli"
	cfg "github.com/candiddev/shared/go/config"
	"github.com/candiddev/shared/go/errs"
	"github.com/candiddev/shared/go/logger"
)

type config struct {
	CLI       cli.Config `json:"cli"`
	CheckCode string     `json:"checkCode"`
	FailWarn  bool       `json:"failWarn"`
	Input     string     `json:"input"`
}

func (c *config) CLIConfig() *cli.Config {
	return &c.CLI
}

func (c *config) Parse(ctx context.Context, configArgs []string, paths string) errs.Err {
	return logger.Error(ctx, cfg.Parse(ctx, c, configArgs, "yaml8n", "", paths))
}
