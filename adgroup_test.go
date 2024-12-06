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
	c.SetHTTPProxy("http://127.0.0.1:7892")
	c.SetHTTPDebug(true)
	req := new(AdGroupListReq)
	req.AdvertiserID = os.Getenv("TikTokAdvertiserID")
	res, err := c.AdGroup.FindAdGroups(req)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res.Data.List[0].IncludedPangleAudiencePackageIDs)
	tmpA := res.Data.List[0].ContextualTagIDs
	fmt.Printf("%#+v\n", len(*tmpA))
}
