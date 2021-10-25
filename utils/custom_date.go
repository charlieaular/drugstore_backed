package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomDate time.Time

var _ json.Unmarshaler = &CustomDate{}

func (cd *CustomDate) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*cd = CustomDate(t)
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(cd).Format("2006-01-02"))
	return []byte(stamp), nil
}
