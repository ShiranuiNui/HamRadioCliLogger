package models

import (
	"encoding/json"
	"strconv"
	"time"
)

type QSO struct {
	MyCallSign         string      `json:"my_call_sign"`
	CallSign           string      `json:"call_sign"`
	Time               ISO8601Time `json:"time"`
	Report             string      `json:"report"`
	Frequency          int         `json:"frequency"`
	Mode               string      `json:"mode"`
	IsRequestedQSLCard bool        `json:"is_requested_qsl_card"`
}

func (q *QSO) ToStrArray() []string {
	return []string{q.MyCallSign, q.CallSign, q.Time.Format("2006-01-02T15:04:05-0700"), q.Report, strconv.Itoa(q.Frequency), q.Mode, strconv.FormatBool(q.IsRequestedQSLCard)}
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
