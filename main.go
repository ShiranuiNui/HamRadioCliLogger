package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/ShiranuiNui/HamRadioCliLogger/models"
	tw "github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "new",
		Usage: "Create New QSO Logging",
		Flags: []cli.Flag{
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
		},
		Action: func(c *cli.Context) error {
			time := models.ISO8601Time{Time: time.Now()}
			qso := models.QSO{MyCallSign: "JM8HBM", CallSign: c.String("callsign"), Time: time, Report: c.String("report"), Frequency: c.Int("freq"), Mode: c.String("mode"), IsRequestedQSLCard: c.Bool("isRequestedQSLCard")}
			table := tw.NewWriter(os.Stdout)
			table.SetAutoFormatHeaders(false)
			table.SetHeader([]string{"MyCallSign", "CallSign", "Time", "Report", "Frequency", "Mode", "IsRequestedQSLCard"})
			table.Append(qso.ToStrArray())
			table.Render()
			json, _ := json.Marshal(qso)
			err := writeJSON(json)
			if err != nil {
				log.Fatal(err)
				return cli.Exit("", 1)
			}
			/*
				qso, err := readJSON()
				if err != nil {
					log.Fatal(err)
					return cli.Exit("", 1)
				}
				var qsoStringArray [][]string
				for _, qso := range qso {
					qsoStringArray = append(qsoStringArray, qso.ToStrArray())
				}
				table := tw.NewWriter(os.Stdout)
				table.SetAutoFormatHeaders(false)
				table.SetHeader([]string{"MyCallSign", "CallSign", "Time", "Report", "Frequency", "Mode", "IsRequestedQSLCard"})
				table.AppendBulk(qsoStringArray)
				table.Render()
			*/
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func writeJSON(json []byte) error {
	usr, _ := user.Current()
	os.Mkdir(usr.HomeDir+"/.hamradio_logger", 0777)
	fp, err := os.OpenFile(usr.HomeDir+"/.hamradio_logger/log.jsonl", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fmt.Fprintln(fp, string(json))
	if err != nil {
		return err
	}
	return nil
}

func readJSON() ([]models.QSO, error) {
	usr, _ := user.Current()
	fp, err := os.OpenFile(usr.HomeDir+"/.hamradio_logger/log.jsonl", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var qso models.QSO
	var qsoArray []models.QSO
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &qso); err != nil {
			return nil, err
		}
		qsoArray = append(qsoArray, qso)
	}
	return qsoArray, nil

}
