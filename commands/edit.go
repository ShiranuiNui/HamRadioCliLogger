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
			Name:    "callsign",
			Usage:   "callsign of worked station",
			Aliases: []string{"c"},
		},
		&cli.StringFlag{
			Name:    "report",
			Usage:   "sent report like rst",
			Aliases: []string{"r"},
		},
		&cli.IntFlag{
			Name:    "freq",
			Usage:   "frequency",
			Aliases: []string{"f"},
		},
		&cli.StringFlag{
			Name:    "mode",
			Usage:   "mode",
			Aliases: []string{"m"},
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
		Name:  "edit-prev",
		Usage: "Edit Previous QSO Logging",
		Flags: flags,
		Action: func(c *cli.Context) error {
			qso, err := repo.ReadPrevQSO(c.String("log_file_path"))
			if err != nil {
				log.Fatal(err)
				return cli.Exit("", 1)
			}
			if c.IsSet("callsign") {
				qso.CallSign = c.String("callsign")
			}
			if c.IsSet("report") {
				qso.Report = c.String("report")
			}
			if c.IsSet("freq") {
				qso.Frequency = c.Int("freq")
			}
			if c.IsSet("mode") {
				qso.Mode = c.String("mode")
			}
			if c.IsSet("isRequestedQSLCard") {
				qso.QSLCardStatuses.IsRequestedQSLCard = c.Bool("isRequestedQSLCard")
			}
			if c.IsSet("qsl_remarks") {
				qso.QSLRemarks = c.String("qsl_remarks")
			}
			if c.IsSet("remarks") {
				qso.Remarks = c.String("remarks")
			}
			if err := repo.EditLatestQSO(qso, c.String("log_file_path")); err != nil {
				log.Fatal(err)
				return cli.Exit("", 1)
			}
			repo.ShowQSOTable([]models.QSO{qso})
			return nil
		}})
}
