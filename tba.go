package tba

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
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
func (c *Client) get(url string, resp interface{}, query ...interface{}) error {
	res, err := c.client.Get(url, query...)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}

// post 处理post请求
// func (c *Client) post(url string, resp interface{}, data ...interface{}) error {
// 	res, err := c.client.Post(url, data...)
// 	if err != nil {
// 		return err
// 	}
// 	return c.rawJson(res, resp)
// }

// structPtr2Map 将结构体指针转换为map
func structPtr2Map(obj interface{}, tagName string) map[string]interface{} {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	return structToMap(v, tagName)
}

// structToMap 将结构体转换为map
func structToMap(v reflect.Value, tagName string) map[string]interface{} {
	t := v.Type()
	data := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 处理匿名嵌套结构体
		if field.Anonymous {
			if value.Kind() == reflect.Struct {
				nestedMap := structToMap(value, tagName)
				for k, v := range nestedMap {
					data[k] = v
				}
			}
			continue
		}

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
			fieldValue = structToMap(value, tagName)
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
