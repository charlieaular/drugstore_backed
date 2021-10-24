package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomDate time.Time

var _ json.Unmarshaler = &CustomDate{}

func (mt *CustomDate) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = CustomDate(t)
	return nil
}

func (mt CustomDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(mt).Format("2006-01-02"))
	return []byte(stamp), nil
}
