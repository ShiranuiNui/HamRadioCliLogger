package models

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

type QSO struct {
	CallSign             string               `json:"call_sign"`
	Time                 ISO8601Time          `json:"time"`
	Report               string               `json:"report"`
	Frequency            int                  `json:"frequency"`
	Mode                 string               `json:"mode"`
	QSLCardStatuses      QSLCardStatuses      `json:"qsl_card_statuses"`
	QSLRemarks           string               `json:"qsl_remarks"`
	Remarks              string               `json:"remarks"`
	MyStationInfomations MyStationInfomations `json:"my_stations_infomations"`
}

type QSLCardStatuses struct {
	IsRequestedQSLCard bool `json:"is_requested_qsl_card"`
	IsSentQSLCard      bool `json:"is_sent_qsl_card"`
	IsReceivedQSLCard  bool `json:"is_reveiced_qsl_card"`
}

type MyStationInfomations struct {
	MyCallSign string `json:"my_call_sign"`
	QTH        string `json:"qth"`
	Rig        string `json:"rig"`
	Output     string `json:"output"`
	Antenna    string `json:"antenna"`
}

func NewQSOFromCliContext(c *cli.Context) QSO {
	time := ISO8601Time{Time: time.Now()}
	qslCardStatuses := QSLCardStatuses{
		IsRequestedQSLCard: c.Bool("isRequestedQSLCard"),
	}
	myStationsInfomations := MyStationInfomations{
		MyCallSign: c.String("mycallsign"),
		QTH:        c.String("qth"),
		Rig:        c.String("rig"),
		Output:     c.String("output"),
		Antenna:    c.String("antenna"),
	}
	qso := QSO{
		CallSign:             c.String("callsign"),
		Time:                 time,
		Report:               c.String("report"),
		Frequency:            c.Int("freq"),
		Mode:                 c.String("mode"),
		QSLCardStatuses:      qslCardStatuses,
		QSLRemarks:           c.String("qsl_remarks"),
		Remarks:              c.String("remarks"),
		MyStationInfomations: myStationsInfomations,
	}
	return qso
}

func (q *QSO) ToStrArray() []string {
	return []string{
		q.CallSign,
		q.Time.Format("2006-01-02T15:04-0700"),
		q.Report, strconv.Itoa(q.Frequency),
		q.Mode,
		q.Remarks,
	}
}

type ISO8601Time struct {
	time.Time
}

func (mt *ISO8601Time) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("\"2006-01-02T15:04:05-0700\"", string(data))
	*mt = ISO8601Time{t}
	return err
}

func (mt ISO8601Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(mt.Format("2006-01-02T15:04:05-0700"))
}
