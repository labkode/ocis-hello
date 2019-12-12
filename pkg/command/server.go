package command

import (
	"github.com/cs3org/reva/cmd/revad/runtime"
	"github.com/gofrs/uuid"
	"github.com/micro/cli"
	"github.com/owncloud/ocis-hello/pkg/config"
	"github.com/owncloud/ocis-hello/pkg/flagset"
	"os"
	"path"
)

// Server is the entrypoint for the server command.
func Server(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start Reva server",
		Flags: flagset.RootWithConfig(cfg),
		Before: func(c *cli.Context) error {
			return nil
		},
		Action: func(c *cli.Context) error {
			//logger := NewLogger(cfg)
			uuid := uuid.Must(uuid.NewV4())
			pidFile := path.Join(os.TempDir(), "revad-"+uuid.String()+".pid")
			runtime.Run(cfg.File, pidFile)
			return nil
		},
	}
}
