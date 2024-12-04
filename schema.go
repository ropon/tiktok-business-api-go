package tba

import (
	"encoding/json"
	"time"
)

const dateFormat = "2006-01-02 15:04:04"

type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(dateFormat))
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}

	*t = UnixTime(time.Unix(timestamp, 0))
	return nil
}

type DateTime time.Time

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(dateFormat))
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var dateStr string

	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}

	parsed, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return err
	}

	*d = DateTime(parsed)
	return nil
}
