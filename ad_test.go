package tba

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run TestFindAds
func TestFindAds(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)
	req := new(AdListReq)
	req.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	//req.Filtering = AdFilter{
	//	AdIDs: []string{""},
	//}

	res, err := c.Ad.FindAds(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res.Data.List)
	//tmpA := res.Data.List[0].ImageIDs
	//fmt.Printf("%#+v\n", len(*tmpA))
}
