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

// go test -v -run TestCreateAd
func TestCreateAd(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	data := new(CreateAdReq)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")

	res, err := c.Ad.CreateAd(data)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(res.Data.AdIDs) == 0 {
		t.Error("Failed to create ad")
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}

// go test -v -run TestUpdateAd$
func TestUpdateAd(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	data := new(UpdateAdReq)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.AdGroupID = ""
	res, err := c.Ad.UpdateAd(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res)
}

// go test -v -run TestUpdateAdStatus
func TestUpdateAdStatus(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	data := new(UpdateAdStatusReq)
	data.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	data.AdIDs = []string{""}
	res, err := c.Ad.UpdateAdStatus(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res)
}
