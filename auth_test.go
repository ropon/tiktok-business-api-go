package tba

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// go test -v -run TestNewTokenConfig
func TestNewTokenConfig(t *testing.T) {
	t.Parallel()

	auth := NewTokenConfig("a", "b")

	assert.Equal(t, "a", auth.AppID)

	auth2 := NewTokenConfig(os.Getenv("TikTokAppID"), os.Getenv("TikTokSecret"))
	if os.Getenv("https_proxy") != "" {
		auth2.SetHTTPProxy(os.Getenv("https_proxy"))
	}
	res, err := auth2.ExchangeAccessToken("")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(res.Data)
}
