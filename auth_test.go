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
	auth2.SetHTTPProxy("http://127.0.0.1:7892")
	res, err := auth2.ExchangeAccessToken("")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(res.Data)
}
