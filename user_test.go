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
	c.SetHTTPProxy("http://127.0.0.1:7892")
	c.SetHTTPDebug(true)

	res, err := c.User.GetUserInfo()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#+v\n", res.Data)
}
