package tba

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/ropon/requests/v2"
)

const (
	defaultBaseURL = "https://business-api.tiktok.com/open_api/v1.3/"
	defaultTimeout = 30 * time.Second
	userAgent      = "tiktok-business-api-go"
)

// Client 客户端
type Client struct {
	client *requests.Request
	common service

	User       *UserService
	Advertiser *AdvertiserService
	Campaign   *CampaignService
	AdGroup    *AdGroupService
	Ad         *AdService
	Report     *ReportService
}

// service 服务
type service struct {
	client *Client
}

// SetDefault 设置默认请求配置
func SetDefault(req *requests.Request) {
	_ = req.SetBaseUrl(defaultBaseURL)
	req.SetTimeout(defaultTimeout)
	req.SetHeader("User-Agent", userAgent)
}

// NewClient 创建客户端
func NewClient(httpReq *requests.Request, accessToken ...string) *Client {
	if httpReq == nil {
		httpReq = requests.New()
		SetDefault(httpReq)
	}
	c := &Client{
		client: httpReq,
	}
	if len(accessToken) > 0 {
		c.client.SetHeader("Access-Token", accessToken[0])
	}
	c.common.client = c

	c.User = (*UserService)(&c.common)
	c.Advertiser = (*AdvertiserService)(&c.common)
	c.Campaign = (*CampaignService)(&c.common)
	c.AdGroup = (*AdGroupService)(&c.common)
	c.Ad = (*AdService)(&c.common)
	c.Report = (*ReportService)(&c.common)
	return c
}

// SetHTTPTimeout 设置http请求超时时间
func (c *Client) SetHTTPTimeout(n time.Duration) {
	c.client.SetTimeout(n)
}

// SetHTTPDebug 设置http请求debug
func (c *Client) SetHTTPDebug(flag bool) {
	c.client.Debug = flag
}

// SetHTTPProxy 设置http请求代理
func (c *Client) SetHTTPProxy(proxyUrl string) {
	c.client.SetProxy(proxyUrl)
}

// rawJson 处理json响应
func (c *Client) rawJson(res *requests.Response, resp interface{}) error {
	if res.Json().Get("code").Int64() != 0 {
		return errors.New(res.Json().Get("message").String())
	}
	err := res.RawJson(resp)
	if err != nil {
		return err
	}
	return nil
}

// get 处理get请求
func (c *Client) get(apiUrl string, resp interface{}, params ...interface{}) error {
	if len(params) > 0 {
		param := params[0]
		// 构建 URL
		u, err := url.Parse(apiUrl)
		if err != nil {
			return err
		}

		query := u.Query()
		err = addParamsToQuery(query, param)
		if err != nil {
			return err
		}

		u.RawQuery = query.Encode()
		apiUrl = u.String()
	}
	res, err := c.client.Get(apiUrl)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}

// post 处理post请求
func (c *Client) post(url string, resp interface{}, data ...interface{}) error {
	if len(data) > 0 {
		postData := data[0]
		bS, err := json.Marshal(postData)
		if err != nil {
			return err
		}
		res, err := c.client.Post(url, string(bS))
		if err != nil {
			return err
		}
		return c.rawJson(res, resp)
	}
	res, err := c.client.Post(url, data...)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}
