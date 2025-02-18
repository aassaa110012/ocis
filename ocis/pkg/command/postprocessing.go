package command

import (
	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/configlog"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/v2/ocis/pkg/command/helper"
	"github.com/owncloud/ocis/v2/ocis/pkg/register"
	"github.com/owncloud/ocis/v2/services/postprocessing/pkg/command"
	"github.com/urfave/cli/v2"
)

// PostprocessingCommand is the entrypoint for the postprocessing command.
func PostprocessingCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     cfg.Postprocessing.Service.Name,
		Usage:    helper.SubcommandDescription(cfg.Postprocessing.Service.Name),
		Category: "services",
		Before: func(c *cli.Context) error {
			configlog.Error(parser.ParseConfig(cfg, true))
			cfg.Postprocessing.Commons = cfg.Commons
			return nil
		},
		Subcommands: command.GetCommands(cfg.Postprocessing),
	}
}

func init() {
	register.AddCommand(PostprocessingCommand)
}
