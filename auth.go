package tba

import (
	"fmt"
	"github.com/ropon/requests/v2"
)

type TokenConfig struct {
	AppID     string
	AppSecret string

	httpReq *requests.Request
}

type BaseResp struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

type TokenInfo struct {
	AccessToken   string   `json:"access_token,omitempty"`
	Scope         []int    `json:"scope,omitempty"`
	AdvertiserIDS []string `json:"advertiser_ids,omitempty"`
}

type TokenResp struct {
	BaseResp
	Data TokenInfo `json:"data,omitempty"`
}

func NewTokenConfig(appID, appSecret string) *TokenConfig {
	return &TokenConfig{
		AppID:     appID,
		AppSecret: appSecret,
		httpReq:   requests.New(),
	}
}

func (t *TokenConfig) SetHTTPDebug(flag bool) {
	t.httpReq.Debug = flag
}

func (t *TokenConfig) SetHTTPProxy(proxyUrl string) {
	t.httpReq.SetProxy(proxyUrl)
}

func (t *TokenConfig) SetDefault(req *requests.Request) error {
	if req == nil {
		req = t.httpReq
	}
	err := req.SetBaseUrl(defaultBaseURL)
	if err != nil {
		return err
	}
	req.SetTimeout(defaultTimeout)
	req.SetHeader("User-Agent", userAgent)
	return nil
}

func (t *TokenConfig) Client(accessToken string) (*requests.Request, error) {
	httpReq := requests.New()
	err := t.SetDefault(httpReq)
	if err != nil {
		return nil, err
	}
	httpReq.SetHeader("Access-Token", accessToken)
	return httpReq, nil
}

func (t *TokenConfig) ExchangeAccessToken(authCode string) (*TokenResp, error) {
	err := t.SetDefault(nil)
	if err != nil {
		return nil, err
	}
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
