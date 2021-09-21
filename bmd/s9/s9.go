package s9

import (
	"github.com/baranovskis/bmd-converter/util"
	"github.com/urfave/cli"
)

var Commands = cli.Command{
	Name:  "s9",
	Usage: "Run season 9 file conversion",
	Subcommands: cli.Commands{
		{
			Name:  "quest_words",
			Usage: "Convert Quest Words file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &QuestWords{})
			},
		},
		{
			Name:  "quest_progress",
			Usage: "Convert Quest Progress file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &QuestProgress{})
			},
		},
		{
			Name:  "npc_dialog",
			Usage: "Convert Npc Dialog file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &NpcDialog{})
			},
		},
		{
			Name:  "npc_name",
			Usage: "Convert Npc Name file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &NpcName{})
			},
		},
		{
			Name:  "buff_effect",
			Usage: "Convert Buff Effect file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &BuffEffect{})
			},
		},
		{
			Name:  "attribute_variation",
			Usage: "Convert Attribute Variation file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &AttributeVariation{})
			},
		},
		{
			Name:  "arca_battle_booty_mix",
			Usage: "Convert Arca Battle Booty Mix file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &ArcaBattleBootyMix{})
			},
		},
		{
			Name:  "item_level_tooltip",
			Usage: "Convert Item Level Tooltip file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &ItemLevelTooltip{})
			},
		},
		{
			Name:  "bonus_exp",
			Usage: "Convert Bonus Exp file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &BonusExp{})
			},
		},
		{
			Name:  "credit",
			Usage: "Convert Credit file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &Credit{})
			},
		},
		{
			Name:  "evolution_monster_box",
			Usage: "Convert Evolution Monster Box file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &EvolutionMonsterBox{})
			},
		},
		{
			Name:  "filter",
			Usage: "Convert Filter file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &Filter{})
			},
		},
		{
			Name:  "filter_name",
			Usage: "Convert Filter Name file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &FilterName{})
			},
		},
		{
			Name:  "gate",
			Usage: "Convert Gate file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &Gate{})
			},
		},
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
			Name:  "item_tooltip",
			Usage: "Convert Item Tooltip file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &ItemTooltip{})
			},
		},
		{
			Name:  "item_tooltip_text",
			Usage: "Convert Item Tooltip Text file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &ItemTooltipText{})
			},
		},
		{
			Name:  "movereq",
			Usage: "Convert Move Req file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &MoveReq{})
			},
		},
		{
			Name:  "minimap",
			Usage: "Convert Mini Map file format",
			Action: func(c *cli.Context) error {
				return util.CommandAction(c, &MiniMap{})
			},
		},
	},
}
