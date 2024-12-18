package tba

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run TestGetAdvertiserInfo
func TestGetAdvertiserInfo(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	res, err := c.Advertiser.GetAdvertiserInfo(&AdvertiserReq{
		AdvertiserIDs: []string{os.Getenv("TikTokAdvertiserID")},
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}
