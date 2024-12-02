package tba

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	c := NewClient(nil)
	assert.Equal(t, defaultBaseURL, c.baseURL.String())
	assert.Equal(t, userAgent, c.UserAgent)

	c2 := NewClient(nil)
	assert.NotSame(t, c.client, c2.client, "NewClient returned same requests.Request, but they should differ")
}

func TestSetHTTPDebug(t *testing.T) {
	t.Parallel()

	client := NewClient(nil)
	client.SetHTTPDebug(true)
	assert.True(t, client.httpDebug)
	client.SetHTTPDebug(false)
	assert.False(t, client.httpDebug)
}
