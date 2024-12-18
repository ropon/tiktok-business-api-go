package tba

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run ListReports
func TestListReports(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)
	req := new(ReportListReq)
	req.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	req.Dimensions = []string{"ad_id", "stat_time_hour"}
	req.ReportType = ReportTypeBasic
	req.DataLevel = DataLevelAuctionAd
	req.StartDate = "2024-10-26"
	req.EndDate = "2024-10-26"
	req.Metrics = []string{"spend", "clicks", "impressions", "reach", "billed_cost"}
	req.EnableTotalMetrics = true
	req.PageSize = 100
	req.Filtering = []ReportFilter{
		{
			FieldName:   "campaign_ids",
			FilterType:  FilterTypeIn,
			FilterValue: "[]",
		},
		{
			FieldName:   "ad_ids",
			FilterType:  FilterTypeIn,
			FilterValue: "[]",
		},
	}
	res, err := c.Report.ListReports(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("total: %#+v\n", res.Data.TotalMetrics)
	fmt.Println("=====")
	fmt.Printf("repot: %#+v\n", res.Data.List)
}
