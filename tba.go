package tba

import (
	"encoding/json"
	"fmt"
	"github.com/ropon/requests/v2"
	"net/url"
	"reflect"
	"strconv"
	"strings"
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
	Campaign   *CampaignService
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
	c.Campaign = (*CampaignService)(&c.common)
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
	res, err := c.client.Get(url, query...)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}

func (c *Client) post(url string, resp interface{}, data ...interface{}) error {
	res, err := c.client.Post(url, data...)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}

func structPtr2Map(obj interface{}, tagName string) map[string]interface{} {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	data := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get(tagName)
		if tag == "-" {
			continue // 忽略标记为 "-" 的字段
		}
		tagParts := strings.Split(tag, ",")
		if len(tagParts) == 0 {
			continue
		}
		key := tagParts[0]
		if key == "" {
			continue
		}

		var fieldValue interface{}
		if value.Kind() == reflect.Ptr && !value.IsNil() {
			fieldValue = structPtr2Map(value.Interface(), tagName)
		} else if value.Kind() == reflect.Struct {
			fieldValue = structPtr2Map(value.Interface(), tagName)
		} else if value.Kind() == reflect.Slice && value.Type().Elem().Kind() == reflect.String {
			// 使用 json.Marshal 处理 []string
			bS, err := json.Marshal(value.Interface())
			if err == nil {
				fieldValue = string(bS)
			}
		} else {
			fieldValue = value.Interface()
		}

		if !value.IsZero() {
			data[key] = fieldValue
		} else {
			// 检查是否有默认值
			for _, part := range tagParts[1:] {
				if strings.HasPrefix(part, "default=") {
					defaultValue := strings.TrimPrefix(part, "default=")
					switch value.Kind() {
					case reflect.Int, reflect.Int64:
						if intValue, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
							data[key] = intValue
						}
					case reflect.String:
						data[key] = defaultValue
						// 可以根据需要添加其他类型的处理
					default:
						panic("unhandled default case")
					}
					break
				}
			}
		}
	}
	return data
}
