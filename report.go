package tba

const (
	ReportTypeBasic            = "BASIC"
	ReportTypeAudience         = "AUDIENCE"
	ReportTypePlayableMaterial = "PLAYABLE_MATERIAL"
	ReportTypeCatalog          = "CATALOG"
	ReportTypeBC               = "BC"
	ReportTypeTTShop           = "TT_SHOP"
)

const (
	DataLevelAuctionAd         = "AUCTION_AD"
	DataLevelAuctionAdGroup    = "AUCTION_ADGROUP"
	DataLevelAuctionCampaign   = "AUCTION_CAMPAIGN"
	DataLevelAuctionAdvertiser = "AUCTION_ADVERTISER"
)

const (
	FilterTypeIn           = "IN"
	FilterTypeMatch        = "MATCH"
	FilterTypeGreaterEqual = "GREATER_EQUAL"
	FilterTypeGreaterThan  = "GREATER_THAN"
	FilterTypeLowerEqual   = "LOWER_EQUAL"
	FilterTypeLowerThan    = "LOWER_THAN"
	FilterTypeBetween      = "BETWEEN"
)

type ReportService service

type Report struct {
	ReportDimensions `json:"dimensions"`
	ReportMetrics    `json:"metrics"`
}

type ReportDimensions struct {
	AdvertiserID       string `json:"advertiser_id"`        // 广告主ID。当dimensions包含 advertiser_id 时返回
	CampaignID         string `json:"campaign_id"`          // 推广系列ID。当dimensions包含 campaign_id 时返回
	AdgroupID          string `json:"adgroup_id"`           // 广告组ID。当dimensions包含 adgroup_id 时返回
	AdID               string `json:"ad_id"`                // 广告ID。当dimensions包含 ad_id 时返回
	StatTimeDay        string `json:"stat_time_day"`        // 消耗发生的时间（天）。格式：2020-01-01 00:00:00
	StatTimeHour       string `json:"stat_time_hour"`       // 消耗发生的时间（小时）。格式：2020-01-01 10:00:00
	AC                 string `json:"ac"`                   // 受众网络类型。详见枚举值-广告管理-网络类型
	Age                string `json:"age"`                  // 受众年龄区间。详见枚举值-广告管理-受众年龄区间
	CountryCode        string `json:"country_code"`         // 受众国家或地区代码。详见附录-地区代码
	InterestCategory   string `json:"interest_category"`    // 旧版一级兴趣分类（将在下个API版本中废弃）
	InterestCategoryV2 string `json:"interest_category_v2"` // 新版一级兴趣分类。使用/tool/interest_category/接口获取完整列表
	Gender             string `json:"gender"`               // 受众性别。枚举值: FEMALE，MALE，NONE
	Language           string `json:"language"`             // 受众语言。详见枚举值-广告管理-受众语言
	Placement          string `json:"placement"`            // 投放版位，详见枚举值-广告管理-版位
	Platform           string `json:"platform"`             // 受众操作系统，详见枚举值-广告管理-受众操作系统
	ContextualTag      string `json:"contextual_tag"`       // 内容相关定向标签
}

type ReportMetrics struct {
	AdvertiserName             string `json:"advertiser_name"`                // 账户名称
	AdvertiserID               string `json:"advertiser_id"`                  // 账户ID
	CampaignName               string `json:"campaign_name"`                  // 推广系列名称
	CampaignID                 string `json:"campaign_id"`                    // 推广系列ID
	ObjectiveType              string `json:"objective_type"`                 // 推广目标
	SplitTest                  string `json:"split_test"`                     // 拆分对比实验状态
	CampaignBudget             string `json:"campaign_budget"`                // 推广系列预算
	CampaignDedicateType       string `json:"campaign_dedicate_type"`         // 推广系列类型
	AppPromotionType           string `json:"app_promotion_type"`             // 应用推广类型
	AdgroupName                string `json:"adgroup_name"`                   // 广告组名称
	AdgroupID                  string `json:"adgroup_id"`                     // 广告组ID
	PlacementType              string `json:"placement_type"`                 // 版位类型
	PromotionType              string `json:"promotion_type"`                 // 推广对象类型
	OptStatus                  string `json:"opt_status"`                     // 程序化创意
	AdgroupDownloadURL         string `json:"adgroup_download_url"`           // 下载URL/网站URL
	ProfileImage               string `json:"profile_image"`                  // 头像
	DPATargetAudienceType      string `json:"dpa_target_audience_type"`       // DPA广告的目标受众类型
	Budget                     string `json:"budget"`                         // 广告组预算
	SmartTarget                string `json:"smart_target"`                   // 优化目标
	BillingEvent               string `json:"billing_event"`                  // 计费点
	BidStrategy                string `json:"bid_strategy"`                   // 竞价策略
	Bid                        string `json:"bid"`                            // 出价
	BidSecondaryGoal           string `json:"bid_secondary_goal"`             // 深层目标出价
	AEOType                    string `json:"aeo_type"`                       // 应用内事件优化类型
	AdName                     string `json:"ad_name"`                        // 广告名称
	AdID                       string `json:"ad_id"`                          // 广告ID
	AdText                     string `json:"ad_text"`                        // 广告标题
	CallToAction               string `json:"call_to_action"`                 // 行动引导文案
	AdProfileImage             string `json:"ad_profile_image"`               // 头像（广告层级）
	AdURL                      string `json:"ad_url"`                         // URL（广告层级）
	TTAppID                    string `json:"tt_app_id"`                      // 推广应用ID
	TTAppName                  string `json:"tt_app_name"`                    // 推广应用名称
	MobileAppID                string `json:"mobile_app_id"`                  // 推广应用在Google Play或Apple App Store中的ID
	ImageMode                  string `json:"image_mode"`                     // 样式
	Currency                   string `json:"currency"`                       // 货币
	IsACO                      bool   `json:"is_aco"`                         // 广告是否为程序化创意广告或智能创意广告
	IsSmartCreative            bool   `json:"is_smart_creative"`              // 广告是否为智能创意广告
	Spend                      string `json:"spend"`                          // 花费：你的广告消耗总金额
	BilledCost                 string `json:"billed_cost"`                    // 净消耗：不包括使用的广告赠款或优惠券的广告消耗总金额
	CashSpend                  string `json:"cash_spend"`                     // 现金消耗：投放广告产生的现金消耗
	VoucherSpend               string `json:"voucher_spend"`                  // 赠款消耗：投放广告产生的赠款消耗
	CPC                        string `json:"cpc"`                            // 平均点击成本（目标页面）：跳转到指定目标页面的每次点击平均成本
	CPM                        string `json:"cpm"`                            // 千次展示成本 (CPM)：每 1,000 次展示平均花费的金额
	Impressions                string `json:"impressions"`                    // 展示量：广告展示的次数
	GrossImpressions           string `json:"gross_impressions"`              // 总展示数（包括无效展示）：广告显示在屏幕上的次数，包括无效展示
	Clicks                     string `json:"clicks"`                         // 点击量（目标页面）：跳转到指定目标页面的点击次数
	CTR                        string `json:"ctr"`                            // 点击率（目标页面）：促成目标页面点击的展示次数占总展示次数的百分比
	Reach                      string `json:"reach"`                          // 覆盖人数：至少看过你的广告一次的去重用户数
	CostPer1000Reached         string `json:"cost_per_1000_reached"`          // 覆盖千人成本：触达 1,000 个去重用户的平均成本
	Frequency                  string `json:"frequency"`                      // 频次：特定时间段内每位用户看到你的广告的平均次数
	Conversion                 string `json:"conversion"`                     // 转化量：广告达成你选择的优化事件的次数
	CostPerConversion          string `json:"cost_per_conversion"`            // 平均转化成本：在转化上花费的平均金额
	ConversionRate             string `json:"conversion_rate"`                // 转化率（点击）：成效数占广告获得的目标页面总点击量的百分比
	ConversionRateV2           string `json:"conversion_rate_v2"`             // 转化率（展示）：成效数占广告获得的总展示次数的百分比
	RealTimeConversion         string `json:"real_time_conversion"`           // 实时转化量：广告达成你选择的优化事件的次数
	RealTimeCostPerConversion  string `json:"real_time_cost_per_conversion"`  // 实时平均转化成本：在转化上花费的平均金额
	RealTimeConversionRate     string `json:"real_time_conversion_rate"`      // 实时转化率（点击）：转化数占广告获得的目标页面总点击量的百分比
	RealTimeConversionRateV2   string `json:"real_time_conversion_rate_v2"`   // 实时转化率（展示）：转化数占广告获得的总展示次数的百分比
	Result                     string `json:"result"`                         // 成效数：广告达到预期的推广目标和优化目标的次数
	CostPerResult              string `json:"cost_per_result"`                // 平均成效成本：广告获得单次成效的平均成本
	ResultRate                 string `json:"result_rate"`                    // 成效率：产生的成效数占广告获得的总展示次数的百分比
	RealTimeResult             string `json:"real_time_result"`               // 实时成效数：广告达到预期的推广目标和优化目标的次数
	RealTimeCostPerResult      string `json:"real_time_cost_per_result"`      // 实时平均成效成本：广告获得单次成效的平均成本
	RealTimeResultRate         string `json:"real_time_result_rate"`          // 实时成效率：产生的成效数占广告获得的总展示次数的百分比
	SecondaryGoalResult        string `json:"secondary_goal_result"`          // 深层目标成效数：广告达到你选择的预期深层目标的次数
	CostPerSecondaryGoalResult string `json:"cost_per_secondary_goal_result"` // 深层目标成效平均成本：广告获得单次深层目标成效的平均成本
	SecondaryGoalResultRate    string `json:"secondary_goal_result_rate"`     // 深层目标成效率：产生的深层目标成效数占广告获得的总展示次数的百分比
}

type ReportData struct {
	TotalMetrics ReportMetrics `json:"total_metrics"`
	List         []Report      `json:"list"`
	PageInfo     PageInfo      `json:"page_info"`
}

type ReportFilter struct {
	FieldName   string `form:"field_name"`   // 筛选字段名称，条件必填
	FilterType  string `form:"filter_type"`  // 筛选类型，条件必填
	FilterValue string `form:"filter_value"` // 要筛选的值，条件必填
}

type ReportListReq struct {
	BaseReq
	AdvertiserIDs           []string     `form:"advertiser_ids"`               // 广告主ID列表，条件必填，最大数量：5
	BCID                    string       `form:"bc_id"`                        // 商务中心ID，条件必填
	ServiceType             string       `form:"service_type"`                 // 广告服务类型，枚举值：AUCTION, RESERVATION
	ReportType              string       `form:"report_type"`                  // 报表类型，必填，枚举值：BASIC, AUDIENCE, PLAYABLE_MATERIAL, CATALOG, BC, TT_SHOP
	DataLevel               string       `form:"data_level"`                   // 报表数据层级，条件必填
	Dimensions              []string     `form:"dimensions"`                   // 维度组合，必填
	Metrics                 []string     `form:"metrics"`                      // 要查询的指标
	EnableTotalMetrics      bool         `form:"enable_total_metrics"`         // 是否开启所请求指标的汇总数据
	StartDate               string       `form:"start_date"`                   // 查询起始日期，条件必填，格式：YYYY-MM-DD
	EndDate                 string       `form:"end_date"`                     // 查询结束日期，条件必填，格式：YYYY-MM-DD
	QueryLifetime           bool         `form:"query_lifetime"`               // 是否请求Lifetime指标
	MultiAdvReportInUTCTime bool         `form:"multi_adv_report_in_utc_time"` // 是否将返回的指标设置为UTC+0时区
	OrderField              string       `form:"order_field"`                  // 排序字段
	OrderType               string       `form:"order_type"`                   // 排序方式，枚举值：ASC, DESC
	Filtering               ReportFilter `json:"filtering"`
}

type ReportListResp struct {
	BaseResp
	Data ReportData `json:"data"`
}

// ListReports 获取报告列表 https://business-api.tiktok.com/portal/docs?id=1740302848100353
func (s *ReportService) ListReports(query *ReportListReq) (*ReportListResp, error) {
	apiUrl := "report/integrated/get/"
	resp := new(ReportListResp)
	err := s.client.get(apiUrl, resp, structPtr2Map(query, "form"))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
