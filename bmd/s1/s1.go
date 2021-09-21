package s1

import (
	"github.com/baranovskis/bmd-converter/util"
	"github.com/urfave/cli"
)

var Commands = cli.Command{
	Name:  "s1",
	Usage: "Run season 1 file conversion",
	Subcommands: cli.Commands{
		{
			Name:  "text",
			Usage: "Convert Text file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &Text{})
			},
		},
		{
			Name:  "item",
			Usage: "Convert Item file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &Item{})
			},
		},
		{
			Name:  "slide",
			Usage: "Convert Slide file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &Slide{})
			},
		},
		{
			Name:  "item_set_option",
			Usage: "Convert Item Set Option file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &ItemSetOption{})
			},
		},
		{
			Name:  "movereq",
			Usage: "Convert Move Req file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &MoveReq{})
			},
		},
	},
}
