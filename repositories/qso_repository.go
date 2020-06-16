package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/ShiranuiNui/HamRadioCliLogger/models"
	tw "github.com/olekukonko/tablewriter"
)

func ReadAllQSO() ([]models.QSO, error) {
	usr, _ := user.Current()
	logFileBytes, err := ioutil.ReadFile(usr.HomeDir + "/.hamradio_logger/log.json")
	if err != nil {
		return nil, err
	}
	if len(logFileBytes) == 0 {
		return []models.QSO{}, nil
	}
	var qsoArray []models.QSO
	if err := json.Unmarshal(logFileBytes, &qsoArray); err != nil {
		return nil, err
	}
	return qsoArray, nil
}

func WriteQSO(qso models.QSO) error {
	usr, _ := user.Current()
	if _, err := os.Stat(usr.HomeDir + "/.hamradio_logger"); os.IsNotExist(err) {
		os.Mkdir(usr.HomeDir+"/.hamradio_logger", 0777)
	}
	fp, err := os.OpenFile(usr.HomeDir+"/.hamradio_logger/log.json", os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()
	currentQSOArray, err := ReadAllQSO()
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(append(currentQSOArray, qso))
	if err != nil {
		return err
	}
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
