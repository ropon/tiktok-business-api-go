package tba

import (
	"fmt"
	"github.com/ropon/requests/v2"
	"net/url"
	"time"
)

const (
	defaultBaseURL = "https://business-api.tiktok.com/open_api/v1.3/"
	defaultTimeout = 30 * time.Second
	userAgent      = "tiktok-business-api-go"
)

type Client struct {
	client    *requests.Request
	baseURL   *url.URL
	UserAgent string
	httpDebug bool

	common service

	User       *UserService
	Advertiser *AdvertiserService
}

type service struct {
	client *Client
}

func NewClient(httpReq *requests.Request) *Client {
	if httpReq == nil {
		httpReq = requests.New()

	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:    httpReq,
		baseURL:   baseURL,
		UserAgent: userAgent,
	}
	c.common.client = c

	c.User = (*UserService)(&c.common)
	c.Advertiser = (*AdvertiserService)(&c.common)
	return c
}

func (c *Client) SetHTTPTimeout(n time.Duration) {
	c.client.SetTimeout(n)
}

func (c *Client) SetHTTPDebug(flag bool) {
	c.client.Debug = flag
}

func (c *Client) SetHTTPProxy(proxyUrl string) {
	c.client.SetProxy(proxyUrl)
}

func (c *Client) rawJson(res *requests.Response, resp interface{}) error {
	if res.Json().Get("code").Int64() != 0 {
		return fmt.Errorf(res.Json().Get("message").String())
	}
	err := res.RawJson(resp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) get(url string, resp interface{}, query ...interface{}) error {
	res, err := c.client.Get(url, query)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}

func (c *Client) post(url string, resp interface{}, data ...interface{}) error {
	res, err := c.client.Post(url, data)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}
