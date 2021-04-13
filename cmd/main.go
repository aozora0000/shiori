package main

import (
	"fmt"
	"github.com/aozora0000/shiori"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:   "shiori",
		Usage:  "cat file from last time final line",
		Action: shiori.GetCommand,
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "initialize related file database",
				Action: shiori.InitCommand,
			},
			{
				Name:   "get",
				Usage:  "cat file and record final line number",
				Action: shiori.GetCommand,
			},
			{
				Name:   "watch",
				Usage:  "watch file and record final line number when updated",
				Action: shiori.WatchCommand,
			},
			{
				Name:   "clear",
				Usage:  "clear final line number cache",
				Action: shiori.ClearCommand,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	os.Exit(0)
}
