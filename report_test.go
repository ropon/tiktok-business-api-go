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
	c.SetHTTPProxy("http://127.0.0.1:7892")
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
			FieldName:   "ad_ids",
			FilterType:  FilterTypeIn,
			FilterValue: "[11111]",
		},
	}
	res, err := c.Report.ListReports(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("total: %#+v\n", res.Data.TotalMetrics)
	fmt.Println("=====")
	fmt.Printf("repot: %#+v\n", res.Data.List[0])
}
