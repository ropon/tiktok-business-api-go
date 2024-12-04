package tba

type UserService service

// UserInfo 结构体包含用户的基本信息
type UserInfo struct {
	DisplayName string   `json:"display_name,omitempty"` // 用户昵称
	Email       string   `json:"email,omitempty"`        // 电子邮箱（注意：电子邮箱已作数据脱敏处理）
	CoreUserID  string   `json:"core_user_id,omitempty"` // 应用开发者在TikTok for Business平台的用户 ID
	CreateTime  UnixTime `json:"create_time,omitempty"`  // 用户创建时间
	AvatarURL   string   `json:"avatar_url,omitempty"`   // 用户头像的URL
}

// UserInfoResp 用户信息响应
type UserInfoResp struct {
	BaseResp
	Data UserInfo `json:"data,omitempty"`
}

// GetUserInfo 获取用户信息 https://business-api.tiktok.com/portal/docs?id=1739665513181185
func (s *UserService) GetUserInfo() (*UserInfoResp, error) {
	apiUrl := "user/info/"
	resp := new(UserInfoResp)
	err := s.client.get(apiUrl, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
