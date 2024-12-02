package tba

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewTokenConfig(t *testing.T) {
	t.Parallel()

	tokenCfg := NewTokenConfig(os.Getenv("TikTokAppID"), os.Getenv("TikTokSecret"))
	tokenCfg.SetHTTPDebug(true)
	req, err := tokenCfg.Client("")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	c := NewClient(req)
	assert.Equal(t, userAgent, c.UserAgent)
	c.SetHTTPDebug(true)
	userInfo, err := c.User.GetUserInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log(userInfo.Data)
}
