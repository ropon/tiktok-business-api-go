package tba

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run TestFindAdGroups
func TestFindAdGroups(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)
	req := new(AdGroupListReq)
	req.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	//req.Filtering = AdGroupFilter{
	//	AdGroupIDs: []string{""},
	//}
	res, err := c.AdGroup.FindAdGroups(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res.Data.List[0])
	//fmt.Printf("%#+v\n", res.Data.List[0].IncludedPangleAudiencePackageIDs)
	//tmpA := res.Data.List[0].ContextualTagIDs
	//fmt.Printf("%#+v\n", len(*tmpA))
}

// go test -v -run TestCreateAdGroup
func TestCreateAdGroup(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	data := new(AdGroup)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.CampaignID = ""
	data.AdGroupName = "TestAdGroup1"
	data.ScheduleType = "SCHEDULE_TYPE_DAILY"

	res, err := c.AdGroup.CreateAdGroup(data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if res.Data.AdGroupID == "" {
		t.Error("Failed to create campaign")
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}

// go test -v -run TestUpdateAdGroup$
func TestUpdateAdGroup(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)
	data := new(AdGroup)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.AdGroupID = ""
	data.AdGroupName = "TestAdGroup1"
	res, err := c.AdGroup.UpdateAdGroup(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if res.Data.AdGroupID == "" {
		t.Error("Failed to update ad group")
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}

// go test -v -run TestUpdateAdGroupStatus
func TestUpdateAdGroupStatus(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)
	data := new(UpdateAdGroupStatusReq)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.AdGroupIDs = []string{""}
	data.OperationStatus = OperationStatusDisable
	res, err := c.AdGroup.UpdateAdGroupStatus(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(res.Data.AdGroupIDs) == 0 {
		t.Error("Failed to update ad group status")
		return
	}
}

// go test -v -run TestUpdateAdGroupBudget
func TestUpdateAdGroupBudget(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)
	data := new(UpdateAdGroupBudgetReq)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.Budget = []Budget{
		{
			AdGroupID: "",
			Budget:    20,
		},
	}
	//data.ScheduledBudget = []ScheduledBudget{
	//	{
	//		AdGroupID:       "1819484713464849",
	//		ScheduledBudget: 20,
	//	},
	//}
	res, err := c.AdGroup.UpdateAdGroupBudget(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if res.Code != 0 {
		t.Error("Failed to update ad group budget")
		return
	}
}
