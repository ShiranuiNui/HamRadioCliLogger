package repositories

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/user"

	"github.com/ShiranuiNui/HamRadioCliLogger/models"
	tw "github.com/olekukonko/tablewriter"
)

func ReadAllQSO() ([]models.QSO, error) {
	return readJSON()
}

func WriteQSO(qso models.QSO) error {
	jsonBytes, err := json.Marshal(qso)
	if err != nil {
		return err
	}
	usr, _ := user.Current()
	if _, err := os.Stat(usr.HomeDir + "/.hamradio_logger"); os.IsNotExist(err) {
		os.Mkdir(usr.HomeDir+"/.hamradio_logger", 0777)
	}
	fp, err := os.OpenFile(usr.HomeDir+"/.hamradio_logger/log.jsonl", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	if _, err = fmt.Fprintln(fp, string(jsonBytes)); err != nil {
		return err
	}
	return nil
}

func ShowQSOTable(qso []models.QSO) {
	var qsoStringArray [][]string
	for _, qso := range qso {
		qsoStringArray = append(qsoStringArray, qso.ToStrArray())
	}
	table := tw.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"MyCallSign", "CallSign", "Time", "Report", "Frequency", "Mode", "IsRequestedQSLCard"})
	table.AppendBulk(qsoStringArray)
	table.Render()
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
