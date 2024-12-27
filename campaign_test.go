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

// go test -v -run TestCreateCampaign
func TestCreateCampaign(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	data := new(Campaign)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.CampaignName = "TestCampaign3"
	data.CampaignType = CampaignTypeRegular
	data.ObjectiveType = ObjectiveAppPromotion
	data.AppPromotionType = AppPromotionTypeInstall
	data.BudgetMode = BudgetModeDay
	data.Budget = 50

	res, err := c.Campaign.CreateCampaign(data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if res.Data.CampaignID == "" {
		t.Error("Failed to create campaign")
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}

// go test -v -run TestUpdateCampaign$
func TestUpdateCampaign(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	data := new(UpdateCampaignReq)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.CampaignID = ""
	data.CampaignName = "TestCampaign3_1"
	data.Budget = 51

	res, err := c.Campaign.UpdateCampaign(data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if res.Data.CampaignID == "" {
		t.Error("Failed to update campaign")
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}

// go test -v -run TestUpdateCampaignStatus
func TestUpdateCampaignStatus(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	data := new(UpdateCampaignStatusReq)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.OperationStatus = OperationStatusEnable
	data.CampaignIDs = []string{""}

	res, err := c.Campaign.UpdateCampaignStatus(data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if res.Data.Status == "" {
		t.Error("Failed to update campaign status")
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}
