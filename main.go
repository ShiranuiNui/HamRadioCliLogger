package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/user"

	"github.com/ShiranuiNui/HamRadioCliLogger/commands"
	"github.com/ShiranuiNui/HamRadioCliLogger/models"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: commands.CmdList,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
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
