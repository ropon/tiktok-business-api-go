package tba

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run TestFindCampaigns
func TestFindCampaigns(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)
	req := new(CampaignListReq)
	req.AdvertiserID = os.Getenv("TikTokAdvertiserID")

	//req.Filtering = CampaignFilter{
	//	CampaignIDs: []string{""},
	//}

	req.Filtering = CampaignFilter{
		CreationFilterStartTime: "",
	}
	res, err := c.Campaign.FindCampaigns(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res.Data.List)
}
