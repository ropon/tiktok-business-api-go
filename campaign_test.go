package tba

import (
	"os"
	"testing"
)

func TestFindCampaigns(t *testing.T) {
	t.Parallel()

	tokenCfg := NewTokenConfig(os.Getenv("TikTokAppID"), os.Getenv("TikTokSecret"))
	tokenCfg.SetHTTPDebug(true)
	req, err := tokenCfg.Client(os.Getenv("TikTokAccessToken"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	c := NewClient(req)
	c.SetHTTPDebug(true)
	resp, err := c.Campaign.FindCampaigns(&CampaignListReq{
		AdvertiserID: "",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(resp.List)
}
