package tba

import (
	"os"
	"testing"
)

// go test -v -run TestGetUserInfo
func TestGetUserInfo(t *testing.T) {
	t.Parallel()
	//httpReq := NewTokenConfig("", "").Client(os.Getenv("TikTokAccessToken"))
	//httpReq.SetProxy("http://127.0.0.1:7892")

	c := NewClient(nil, os.Getenv("TikTokAccessToken"))
	c.SetHTTPProxy("http://127.0.0.1:7892")

	res, err := c.User.GetUserInfo()
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(res.Data)
}
