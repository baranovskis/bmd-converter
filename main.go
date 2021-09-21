package main

import (
	"log"
	"os"
	"sort"

	"github.com/baranovskis/bmd-converter/bmd/s1"
	"github.com/baranovskis/bmd-converter/bmd/s9"
	"github.com/urfave/cli"
)

const VERSION = "1.5"

func main() {
	app := &cli.App{
		Version: VERSION,
		Name:    "bmd-converter",
		Usage:   "tool for .bmd file conversion",
		UsageText: "bmd-converter.exe [global options] command [command options] [arguments...]\n" +
			"\t * bmd-converter.exe <client season> <bmd type> <file path>\n" +
			"\t * bmd-converter.exe s1 text text.bmd - decode file to .csv\n" +
			"\t * bmd-converter.exe s1 text text.csv - encode file to .bmd",
		Authors: []cli.Author{
			{
				Name:  "NexT",
				Email: "info@baranovskis.dev",
			},
		},
		Commands: []cli.Command{
			s1.Commands,
			s9.Commands,
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
