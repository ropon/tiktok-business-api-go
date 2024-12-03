package tba

import (
	"github.com/ropon/requests/v2"
)

type AdvertiserService service

type AdvertiserInfo struct {
	AdvertiserID          string   `json:"advertiser_id,omitempty"`
	OwnerBcID             string   `json:"owner_bc_id,omitempty"`
	Status                string   `json:"status,omitempty"`
	Role                  string   `json:"role,omitempty"`
	RejectionReason       string   `json:"rejection_reason,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Timezone              string   `json:"timezone,omitempty"`
	DisplayTimezone       string   `json:"display_timezone,omitempty"`
	Company               string   `json:"company,omitempty"`
	CompanyNameEditable   bool     `json:"company_name_editable,omitempty"`
	Industry              string   `json:"industry,omitempty"`
	Address               string   `json:"address,omitempty"`
	Country               string   `json:"country,omitempty"`
	AdvertiserAccountType string   `json:"advertiser_account_type,omitempty"`
	Currency              string   `json:"currency,omitempty"`
	Contacter             string   `json:"contacter,omitempty"`
	Email                 string   `json:"email,omitempty"`
	CellphoneNumber       string   `json:"cellphone_number,omitempty"`
	TelephoneNumber       string   `json:"telephone_number,omitempty"`
	Language              string   `json:"language,omitempty"`
	LicenseNo             string   `json:"license_no,omitempty"`
	LicenseUrl            string   `json:"license_url,omitempty"`
	Description           string   `json:"description,omitempty"`
	Balance               float64  `json:"balance,omitempty"`
	CreateTime            UnixTime `json:"create_time"`
	AdvertiserName        string   `json:"advertiser_name,omitempty"`
}

type AdvertisersReq struct {
	AppID  string `json:"app_id,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type AdvertiserData struct {
	List []AdvertiserInfo `json:"list,omitempty"`
}

type AdvertisersResp struct {
	BaseResp
	Data AdvertiserData `json:"data"`
}

type AdvertiserReq struct {
	AdvertiserIDS []string `json:"advertiser_ids,omitempty"`
	Fields        []string `json:"fields,omitempty"`
}

type AdvertiserResp struct {
	BaseResp
	Data AdvertiserData `json:"data"`
}

func (s *AdvertiserService) GetAdvertisers(query *AdvertisersReq) (*AdvertisersResp, error) {
	apiUrl := "oauth2/advertiser/get/"
	resp := new(AdvertisersResp)
	err := s.client.get(apiUrl, resp, requests.StructPtr2Map(query, "json"))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *AdvertiserService) GetAdvertiserInfo(query *AdvertiserReq) (*AdvertiserResp, error) {
	apiUrl := "advertiser/info/"
	resp := new(AdvertiserResp)
	err := s.client.get(apiUrl, resp, structPtr2Map(query, "json"))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
