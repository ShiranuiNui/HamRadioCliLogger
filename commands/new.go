package commands

import (
	"log"
	"time"

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
	}
	CmdList = append(CmdList, &cli.Command{
		Name:  "new",
		Usage: "Create New QSO Logging",
		Flags: flags,
		Action: func(c *cli.Context) error {
			time := models.ISO8601Time{Time: time.Now()}
			qso := models.QSO{MyCallSign: c.String("mycallsign"), CallSign: c.String("callsign"), Time: time, Report: c.String("report"), Frequency: c.Int("freq"), Mode: c.String("mode"), IsRequestedQSLCard: c.Bool("isRequestedQSLCard")}
			if err := repo.WriteQSO(qso); err != nil {
				log.Fatal(err)
				return cli.Exit("", 1)
			}
			repo.ShowQSOTable([]models.QSO{qso})
			return nil
		}})
}
