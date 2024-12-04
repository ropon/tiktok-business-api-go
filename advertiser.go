package tba

type AdvertiserService service

// AdvertiserStatus 定义广告主状态
type AdvertiserStatus string

const (
	StatusDisable              AdvertiserStatus = "STATUS_DISABLE"                // 该广告账户已关户
	StatusPendingConfirm       AdvertiserStatus = "STATUS_PENDING_CONFIRM"        // 申请待审核
	StatusPendingVerified      AdvertiserStatus = "STATUS_PENDING_VERIFIED"       // 待验证
	StatusConfirmFail          AdvertiserStatus = "STATUS_CONFIRM_FAIL"           // 审核不通过
	StatusEnable               AdvertiserStatus = "STATUS_ENABLE"                 // 审核通过
	StatusConfirmFailEnd       AdvertiserStatus = "STATUS_CONFIRM_FAIL_END"       // CRM审核不通过
	StatusPendingConfirmModify AdvertiserStatus = "STATUS_PENDING_CONFIRM_MODIFY" // 修改待审核
	StatusConfirmModifyFail    AdvertiserStatus = "STATUS_CONFIRM_MODIFY_FAIL"    // 修改审核不通过
	StatusLimit                AdvertiserStatus = "STATUS_LIMIT"                  // 用户被惩罚
	StatusWaitForBpmAudit      AdvertiserStatus = "STATUS_WAIT_FOR_BPM_AUDIT"     // 等待CRM审核
	StatusWaitForPublicAuth    AdvertiserStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"   // 待对公验证
	StatusSelfServiceUnaudited AdvertiserStatus = "STATUS_SELF_SERVICE_UNAUDITED" // 自助开户待验证资质
	StatusContractPending      AdvertiserStatus = "STATUS_CONTRACT_PENDING"       // 合同未生效
)

// AdvertiserRole 定义广告主角色
type AdvertiserRole string

const (
	RoleAdvertiser      AdvertiserRole = "ROLE_ADVERTISER"       // 普通广告主（直客）
	RoleChildAdvertiser AdvertiserRole = "ROLE_CHILD_ADVERTISER" // 普通广告主（代理商子客户）
	RoleChildAgent      AdvertiserRole = "ROLE_CHILD_AGENT"      // 二级代理商
	RoleAgent           AdvertiserRole = "ROLE_AGENT"            // 一级代理商
)

// AdvertiserAccountType 定义广告账户类型
type AdvertiserAccountType string

const (
	AccountTypeReservation AdvertiserAccountType = "RESERVATION" // 合约广告账户
	AccountTypeAuction     AdvertiserAccountType = "AUCTION"     // 竞价广告账户
)

// AdvertiserInfo 定义广告主信息
type AdvertiserInfo struct {
	AdvertiserID            string                `json:"advertiser_id,omitempty"`             // 广告账号 ID
	OwnerBcID               string                `json:"owner_bc_id,omitempty"`               // 广告账户所属的商务中心的 ID（注意：仅在特定条件下返回）
	Status                  AdvertiserStatus      `json:"status,omitempty"`                    // 广告账号状态（枚举值详见枚举值-广告主状态）
	Role                    AdvertiserRole        `json:"role,omitempty"`                      // 广告账号角色（枚举值详见枚举值-广告主角色）
	RejectionReason         string                `json:"rejection_reason,omitempty"`          // 审核拒绝原因
	Name                    string                `json:"name,omitempty"`                      // 广告账号名称
	Timezone                string                `json:"timezone,omitempty"`                  // 广告账号时区信息（如 "Etc/GMT"，"Europe/London"）
	DisplayTimezone         string                `json:"display_timezone,omitempty"`          // 广告账号所在时区名称（如 "Europe/London"）
	Company                 string                `json:"company,omitempty"`                   // 广告账号公司名
	CompanyNameEditable     bool                  `json:"company_name_editable,omitempty"`     // 公司名称是否支持通过 API 修改
	Industry                string                `json:"industry,omitempty"`                  // 广告账号行业类别代码
	Address                 string                `json:"address,omitempty"`                   // 广告账号地址信息
	Country                 string                `json:"country,omitempty"`                   // 广告账号注册地代码
	AdvertiserAccountType   AdvertiserAccountType `json:"advertiser_account_type,omitempty"`   // 广告账户类型（RESERVATION 或 AUCTION）
	Currency                string                `json:"currency,omitempty"`                  // 广告账号使用的货币类型（ISO 4217 代码）
	Contacter               string                `json:"contacter,omitempty"`                 // 联系人姓名（已脱敏）
	Email                   string                `json:"email,omitempty"`                     // 广告账号联系人邮箱（已脱敏）
	CellphoneNumber         string                `json:"cellphone_number,omitempty"`          // 联系手机号码（已脱敏）
	TelephoneNumber         string                `json:"telephone_number,omitempty"`          // 固定电话号码（已脱敏）
	Language                string                `json:"language,omitempty"`                  // 广告账号所使用的语言代码
	LicenseNo               string                `json:"license_no,omitempty"`                // 营业执照编号
	LicenseUrl              string                `json:"license_url,omitempty"`               // 营业执照预览 URL（1小时内有效）
	LicenseProvince         string                `json:"license_province,omitempty"`          // 营业执照的颁发省份（即将废弃）
	LicenseCity             string                `json:"license_city,omitempty"`              // 营业执照的颁发城市（即将废弃）
	PromotionArea           string                `json:"promotion_area,omitempty"`            // 广告账户主要推广地域（即将废弃）
	PromotionCenterProvince string                `json:"promotion_center_province,omitempty"` // 广告账户主要推广省份（即将废弃）
	PromotionCenterCity     string                `json:"promotion_center_city,omitempty"`     // 广告账户主要推广城市（即将废弃）
	Brand                   string                `json:"brand,omitempty"`                     // 品牌名称（即将废弃）
	Description             string                `json:"description,omitempty"`               // 品牌描述
	Balance                 float64               `json:"balance,omitempty"`                   // 广告账户可用余额
	CreateTime              UnixTime              `json:"create_time,omitempty"`               // 广告账号创建时间（Unix时间戳，单位秒）
}

// AdvertiserData 广告主列表数据
type AdvertiserData struct {
	List []AdvertiserInfo `json:"list,omitempty"`
}

// AdvertiserReq 定义获取广告主详细信息的请求结构
type AdvertiserReq struct {
	AdvertiserIDs []string `form:"advertiser_ids"` // 需要查看详细信息的广告主 ID 列表
	Fields        []string `form:"fields"`         // 需要获取的字段
}

// AdvertiserResp 获取广告主信息响应
type AdvertiserResp struct {
	BaseResp
	Data AdvertiserData `json:"data,omitempty"`
}

// GetAdvertiserInfo 获取广告主信息 https://business-api.tiktok.com/portal/docs?id=1739593083610113
func (s *AdvertiserService) GetAdvertiserInfo(query *AdvertiserReq) (*AdvertiserResp, error) {
	apiUrl := "advertiser/info/"
	resp := new(AdvertiserResp)
	err := s.client.get(apiUrl, resp, structPtr2Map(query, "form"))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
