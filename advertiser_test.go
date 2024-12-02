package tba

import (
	"os"
	"testing"
)

func TestGetAdvertiserInfo(t *testing.T) {
	t.Parallel()

	tokenCfg := NewTokenConfig(os.Getenv("TikTokAppID"), os.Getenv("TikTokSecret"))
	tokenCfg.SetHTTPDebug(true)
	req, err := tokenCfg.Client("")
	if err != nil {
		t.Error(err.Error())
		return
	}
	c := NewClient(req)
	resp, err := c.Advertiser.GetAdvertiserInfo(&AdvertiserReq{
		AdvertiserIDS: []string{""},
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(resp.Data)

}
