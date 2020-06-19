package commands

import (
	"log"

	"github.com/ShiranuiNui/HamRadioCliLogger/models"
	repo "github.com/ShiranuiNui/HamRadioCliLogger/repositories"
	"github.com/urfave/cli/v2"
)

func init() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "callsign",
			Usage:    "callsign of worked station",
			Aliases:  []string{"c"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "report",
			Usage:    "sent report like rst",
			Aliases:  []string{"r"},
			Required: true,
		},
		&cli.IntFlag{
			Name:     "freq",
			Usage:    "frequency",
			Aliases:  []string{"f"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "mode",
			Usage:    "mode",
			Aliases:  []string{"m"},
			Required: true,
		},
		&cli.BoolFlag{
			Name:    "isRequestedQSLCard",
			Usage:   "Worked Station needs QSL Card?",
			Value:   false,
			Aliases: []string{"q"},
		},
		&cli.StringFlag{
			Name:    "qsl_remarks",
			Usage:   "Remakrs for QSL Card",
			Aliases: []string{"qrmks"},
		},
		&cli.StringFlag{
			Name:    "remarks",
			Usage:   "Remakrs(Not Use for QSL Card)",
			Aliases: []string{"rmks"},
		},
	}
	CmdList = append(CmdList, &cli.Command{
		Name:  "new",
		Usage: "Create New QSO Logging",
		Flags: flags,
		Action: func(c *cli.Context) error {
			qso := models.NewQSOFromCliContext(c)
			if err := repo.WriteQSO(qso, c.String("log_file_path")); err != nil {
				log.Fatal(err)
				return cli.Exit("", 1)
			}
			repo.ShowQSOTable([]models.QSO{qso})
			return nil
		}})
}
