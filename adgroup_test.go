package tba

import (
	"os"
	"testing"
)

// go test -v -run TestFindAdGroups
func TestFindAdGroups(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	c.SetHTTPProxy("http://127.0.0.1:7892")
	c.SetHTTPDebug(true)
	req := new(AdGroupListReq)
	req.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	res, err := c.AdGroup.FindAdGroups(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(res.Data)
}
