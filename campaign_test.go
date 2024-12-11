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
	c.SetHTTPProxy("http://127.0.0.1:7892")
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
