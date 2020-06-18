package commands

import (
	"log"

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
	}
	CmdList = append(CmdList, &cli.Command{
		Name:  "search",
		Usage: "Search QSO Logging by Callsign",
		Flags: flags,
		Action: func(c *cli.Context) error {
			qso, err := repo.ReadQSOByCallsign(c.String("callsign"), c.String("log_file_path"))
			if err != nil {
				log.Fatal(err)
				return cli.Exit("", 1)
			}
			repo.ShowQSOTable(qso)
			return nil
		}})
}
