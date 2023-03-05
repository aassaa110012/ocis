package command

import (
	"context"
	"fmt"

	"github.com/oklog/run"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/configlog"
	"github.com/owncloud/ocis/v2/ocis-pkg/service/grpc"
	"github.com/owncloud/ocis/v2/ocis-pkg/version"
	svcProtogen "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/policies/v0"
	"github.com/owncloud/ocis/v2/services/policies/pkg/authz"
	"github.com/owncloud/ocis/v2/services/policies/pkg/config"
	"github.com/owncloud/ocis/v2/services/policies/pkg/config/parser"
	svcEvent "github.com/owncloud/ocis/v2/services/policies/pkg/service/event"
	svcGRPC "github.com/owncloud/ocis/v2/services/policies/pkg/service/grpc"
	"github.com/urfave/cli/v2"
)

// Server is the entrypoint for the server command.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start the %s service without runtime (unsupervised mode)", "authz"),
		Category: "server",
		Before: func(c *cli.Context) error {
			return configlog.ReturnFatal(parser.ParseConfig(cfg))
		},
		Action: func(c *cli.Context) error {
			var (
				gr          = run.Group{}
				ctx, cancel = func() (context.Context, context.CancelFunc) {
					if cfg.Context == nil {
						return context.WithCancel(context.Background())
					}
					return context.WithCancel(cfg.Context)
				}()
				authorizers []authz.Authorizer
			)
			defer cancel()

			if cfg.OPA.Enabled {
				if opaAuthorizer, err := authz.NewOPA(cfg); err == nil {
					authorizers = append(authorizers, opaAuthorizer)
				} else {
					return err
				}
			}

			{
				svc, err := grpc.NewService(
					grpc.Name(cfg.Service.Name),
					grpc.Namespace(cfg.GRPC.Namespace),
					grpc.Version(version.GetString()),
					grpc.Address(cfg.GRPC.Addr),
					grpc.Context(ctx),
				)
				if err != nil {
					return err
				}

				grpcSvc, err := svcGRPC.New(authorizers)
				if err != nil {
					return err
				}

				if err := svcProtogen.RegisterPoliciesProviderHandler(
					svc.Server(),
					grpcSvc,
				); err != nil {
					return err
				}

				gr.Add(svc.Run, func(_ error) {
					cancel()
				})
			}

			{
				eventSvc, err := svcEvent.New(cfg, authorizers)
				if err != nil {
					return err
				}

				gr.Add(eventSvc.Run, func(_ error) {
					cancel()
				})
			}

			return gr.Run()
		},
	}
}
