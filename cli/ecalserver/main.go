package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/mariusor/esports-calendar/cmd"
)

var version = "(unknown)"

func main() {
	var err error

	ctl := cli.App{
		Name:    "ecalsrv",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "Set storage path",
				Value: "./",
			},
		},
		Commands: []cli.Command{
			cmd.Server,
		},
	}

	err = ctl.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
