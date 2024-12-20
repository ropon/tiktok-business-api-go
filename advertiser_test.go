package tba

import (
	"os"
	"testing"
)

// go test -v -run TestGetAdvertiserInfo
func TestGetAdvertiserInfo(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	c.SetHTTPProxy("http://127.0.0.1:7892")

	res, err := c.Advertiser.GetAdvertiserInfo(&AdvertiserReq{
		AdvertiserIDs: []string{os.Getenv("TikTokAdvertiserID")},
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(res.Data)
}
