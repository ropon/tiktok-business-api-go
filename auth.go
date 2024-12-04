package tba

import (
	"fmt"

	"github.com/ropon/requests/v2"
)

// TokenConfig 获取token配置
type TokenConfig struct {
	AppID     string            // 应用ID
	AppSecret string            // 应用密钥
	httpReq   *requests.Request // 请求
}

// BaseResp 结构体表示API响应的基本结构
type BaseResp struct {
	Code      int    `json:"code,omitempty"`       // 返回码（完整返回码列表及描述，可查看附录-返回码）
	Message   string `json:"message,omitempty"`    // 返回信息（更多信息可查看附录-返回码）
	RequestID string `json:"request_id,omitempty"` // 请求的日志ID，唯一标识一个请求
}

// TokenInfo token信息
type TokenInfo struct {
	AccessToken   string   `json:"access_token,omitempty"`   // 访问令牌
	Scope         []int    `json:"scope,omitempty"`          // 权限范围
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"` // 广告主ID列表
}

// TokenResp token请求响应
type TokenResp struct {
	BaseResp
	Data TokenInfo `json:"data,omitempty"`
}

// NewTokenConfig 创建token配置
func NewTokenConfig(appID, appSecret string) *TokenConfig {
	return &TokenConfig{
		AppID:     appID,
		AppSecret: appSecret,
		httpReq:   requests.New(),
	}
}

// SetHTTPDebug 设置http请求debug
func (t *TokenConfig) SetHTTPDebug(flag bool) {
	t.httpReq.Debug = flag
}

// SetHTTPProxy 设置http请求代理
func (t *TokenConfig) SetHTTPProxy(proxyUrl string) {
	t.httpReq.SetProxy(proxyUrl)
}

// ExchangeAccessToken 交换访问令牌 https://business-api.tiktok.com/portal/docs?id=1738373164380162
func (t *TokenConfig) ExchangeAccessToken(authCode string) (*TokenResp, error) {
	SetDefault(t.httpReq)
	apiUrl := "oauth2/access_token/"
	jsonData := fmt.Sprintf(`{
	"app_id": "%s",
	"secret": "%s",
	"auth_code": "%s"
}`, t.AppID, t.AppSecret, authCode)
	res, err := t.httpReq.Post(apiUrl, jsonData)
	if err != nil {
		return nil, err
	}
	result := new(TokenResp)
	err = res.RawJson(result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, fmt.Errorf(result.Message)
	}
	return result, nil
}

func (t *TokenConfig) Client(accessToken string) *requests.Request {
	httpReq := requests.New()
	SetDefault(httpReq)
	httpReq.SetHeader("Access-Token", accessToken)
	return httpReq
}
