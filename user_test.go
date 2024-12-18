package tba

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run TestGetUserInfo
func TestGetUserInfo(t *testing.T) {
	t.Parallel()

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	if os.Getenv("https_proxy") != "" {
		c.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	c.SetHTTPDebug(true)

	res, err := c.User.GetUserInfo()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}
