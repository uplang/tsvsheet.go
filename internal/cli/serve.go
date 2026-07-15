package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

// serveConfig binds the serve command's spreadsheet path and bind address.
type serveConfig struct {
	source sourcePath
	host   string
	port   int
}

// serveCommand builds the `serve` command.
func serveCommand() *cli.Command {
	cfg := serveConfig{}
	return &cli.Command{
		Name:      cmdServe,
		Usage:     "Serve a browser spreadsheet editor.",
		ArgsUsage: "<sheet>",
		Description: `Host a local web spreadsheet backed by the tsvsheet engine: edit any cell
(a value or an =formula) in the browser, recompute live, and save. The sheet
is a required positional file path (serve saves edits back to it, so stdin is
not accepted).

Examples:
  tsvsheet serve sheet.tsvt
  tsvsheet serve --host 0.0.0.0 --port 8080 sheet.tsvt`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Sources:     cli.EnvVars("HOST"),
				Value:       "127.0.0.1",
				Usage:       "Host address to bind",
				Destination: &cfg.host,
			},
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Sources:     cli.EnvVars("PORT"),
				Value:       8080,
				Usage:       "Port to listen on",
				Destination: &cfg.port,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			cfg.source = positional(c.Args().Slice()).at(0)
			return runServe(ctx, cfg)
		},
	}
}
