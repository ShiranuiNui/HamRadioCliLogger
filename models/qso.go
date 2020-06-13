package models

import (
	"encoding/json"
	"strconv"
	"time"
)

// QSO is nya-n
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

// 新しい構造体を宣言します。元の構造体と同じように扱いたいので embedded してます
type ISO8601Time struct {
	time.Time
}

// Unmarshal時の動作を定義します
func (mt *ISO8601Time) UnmarshalJSON(data []byte) error {
	// 要素がそのまま渡ってくるので "(ダブルクォート)でQuoteされてます
	t, err := time.Parse("\"2006-01-02T15:04:05-0700\"", string(data))
	*mt = ISO8601Time{t}
	return err
}

// Marshal時の動作を定義します
func (mt ISO8601Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(mt.Format("2006-01-02T15:04:05-0700"))
}
