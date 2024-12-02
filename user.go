package tba

type UserService service

type UserInfo struct {
	DisplayName string `json:"display_name,omitempty"`
	Email       string `json:"email,omitempty"`
	CoreUserID  string `json:"core_user_id,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
}

type UserInfoResp struct {
	BaseResp
	Data UserInfo `json:"data,omitempty"`
}

func (s *UserService) GetUserInfo() (*UserInfoResp, error) {
	apiUrl := "user/info/"
	resp := new(UserInfoResp)
	err := s.client.get(apiUrl, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
