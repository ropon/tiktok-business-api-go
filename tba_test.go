package tba

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -run TestNewClient
func TestNewClient(t *testing.T) {
	t.Parallel()

	c := NewClient(nil)

	assert.Equal(t, defaultBaseURL, c.client.BaseUrl.String())

	assert.False(t, c.client.Debug)
	c.SetHTTPDebug(true)
	assert.True(t, c.client.Debug)
}

// go test -v -run TestStructToMap
func TestStructToMap(t *testing.T) {
	t.Parallel()
	req := new(ReportListReq)
	req.Filtering = []ReportFilter{
		{
			FieldName:   "ad_ids",
			FilterType:  FilterTypeIn,
			FilterValue: "[11111]",
		},
	}
	fmt.Printf("%+#v\n", addParamsToQuery(url.Values{}, req))
}
