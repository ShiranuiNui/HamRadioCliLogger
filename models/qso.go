package models

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

type QSO struct {
	MyCallSign         string      `json:"my_call_sign"`
	CallSign           string      `json:"call_sign"`
	Time               ISO8601Time `json:"time"`
	Report             string      `json:"report"`
	Frequency          int         `json:"frequency"`
	Mode               string      `json:"mode"`
	IsRequestedQSLCard bool        `json:"is_requested_qsl_card"`
	QSLRemarks         string      `json:"qsl_remarks"`
	Remarks            string      `json:"remarks"`
}

func NewQSOFromCliContext(c *cli.Context) QSO {
	time := ISO8601Time{Time: time.Now()}
	qso := QSO{MyCallSign: c.String("mycallsign"),
		CallSign:           c.String("callsign"),
		Time:               time,
		Report:             c.String("report"),
		Frequency:          c.Int("freq"),
		Mode:               c.String("mode"),
		IsRequestedQSLCard: c.Bool("isRequestedQSLCard"),
		QSLRemarks:         c.String("qsl_remarks"),
		Remarks:            c.String("remarks"),
	}
	return qso
}

func (q *QSO) ToStrArray() []string {
	return []string{q.MyCallSign,
		q.CallSign,
		q.Time.Format("2006-01-02T15:04:05-0700"),
		q.Report, strconv.Itoa(q.Frequency),
		q.Mode,
		strconv.FormatBool(q.IsRequestedQSLCard),
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
