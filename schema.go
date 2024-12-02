package tba

import (
	"encoding/json"
	"time"
)

const dataFormat = "2006-01-02 15:04:04"

type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(dataFormat))
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	// 解析 JSON 数字
	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}

	// 将 Unix 时间戳转换为 time.Time
	*t = UnixTime(time.Unix(timestamp, 0))
	return nil
}
