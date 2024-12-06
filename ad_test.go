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
	c.SetHTTPProxy("http://127.0.0.1:7892")
	c.SetHTTPDebug(true)
	req := new(AdListReq)
	req.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	res, err := c.Ad.FindAds(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res.Data.List[0].TrackingOfflineEventSetIDs)
	tmpA := res.Data.List[0].ImageIDs
	fmt.Printf("%#+v\n", len(*tmpA))
}
