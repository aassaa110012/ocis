package command

import (
	"github.com/owncloud/ocis/ocis-pkg/config"
	"github.com/owncloud/ocis/ocis/pkg/register"
	"github.com/owncloud/ocis/settings/pkg/command"
	"github.com/urfave/cli/v2"
)

// SettingsCommand is the entry point for the settings command.
func SettingsCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "settings",
		Usage:    "Start settings server",
		Category: "Extensions",
		Before: func(ctx *cli.Context) error {
			if err := ParseConfig(ctx, cfg); err != nil {
				return err
			}

			if cfg.Commons != nil {
				cfg.Settings.Commons = cfg.Commons
			}

			return nil
		},
		Subcommands: command.GetCommands(cfg.Settings),
	}
}

func init() {
	register.AddCommand(SettingsCommand)
}
