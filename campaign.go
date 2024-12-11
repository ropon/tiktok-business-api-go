package tba

type CampaignService service

// CampaignSystemOrigin 推广系列来源
type CampaignSystemOrigin string

const (
	CampaignSystemOriginPromote       CampaignSystemOrigin = "PROMOTE"         //该推广系列为通过 TikTok 移动应用创建的内容加热推广系列
	CampaignSystemOriginTTAdsPlatform CampaignSystemOrigin = "TT_ADS_PLATFORM" //该推广系列为通过 API 或TikTok 广告管理平台创建的非内容加热推广系列
)

// CampaignType 推广系列类型
type CampaignType string

const (
	CampaignTypeRegular CampaignType = "REGULAR_CAMPAIGN" // 普通推广系列
	CampaignTypeIOS14   CampaignType = "IOS14_CAMPAIGN"   // iOS 14专属推广系列
)

// PostbackWindowMode SKAN 4.0 回传模式
type PostbackWindowMode string

const (
	PostbackWindowMode1 PostbackWindowMode = "POSTBACK_WINDOW_MODE1" // 确保能接收到第一次回传，对应的归因窗口期为 0-2 天
	PostbackWindowMode2 PostbackWindowMode = "POSTBACK_WINDOW_MODE2" // 确保能接收到第一次及第二次回传，对应的归因窗口期为 3-7 天
	PostbackWindowMode3 PostbackWindowMode = "POSTBACK_WINDOW_MODE3" // 确保能接收到三次回传，对应的归因窗口期为 8-35 天
)

// BudgetMode 预算类型
type BudgetMode string

const (
	BudgetModeInfinite     BudgetMode = "BUDGET_MODE_INFINITE"             // 不限预算
	BudgetModeTotal        BudgetMode = "BUDGET_MODE_TOTAL"                // 总预算
	BudgetModeDay          BudgetMode = "BUDGET_MODE_DAY"                  // 日预算
	BudgetModeDynamicDaily BudgetMode = "BUDGET_MODE_DYNAMIC_DAILY_BUDGET" // 动态日预算
)

// OperationStatus 推广系列的操作状态
type OperationStatus string

const (
	OperationStatusEnable  OperationStatus = "ENABLE"  // 推广系列处于启用（"开"）状态
	OperationStatusDisable OperationStatus = "DISABLE" // 推广系列处于未启用（"关"）状态
)

// Objective 推广类型
type Objective string

const (
	ObjectiveApp         Objective = "APP"          // 应用
	ObjectiveLandingPage Objective = "LANDING_PAGE" // 落地页
)

// AppPromotionType 应用推广类型
type AppPromotionType string

const (
	AppPromotionTypeInstall         AppPromotionType = "APP_INSTALL"         // 应用安装
	AppPromotionTypeRetargeting     AppPromotionType = "APP_RETARGETING"     // 应用再营销
	AppPromotionTypePreregistration AppPromotionType = "APP_PREREGISTRATION" // 应用预注册
	AppPromotionTypePostsPromotion  AppPromotionType = "APP_POSTS_PROMOTION" // 应用帖子推广
)

// CampaignProductSource 推广系列的商品来源
type CampaignProductSource string

const (
	CampaignProductSourceCatalog CampaignProductSource = "CATALOG" // 商品库
	CampaignProductSourceStore   CampaignProductSource = "STORE"   // TikTok Shop 店铺，或 TikTok 橱窗
)

// CampaignAppProfilePageState 下载中间页使用情况
type CampaignAppProfilePageState string

const (
	CampaignAppProfilePageStateInvalid CampaignAppProfilePageState = "INVALID"
	CampaignAppProfilePageStateUnset   CampaignAppProfilePageState = "UNSET"
	CampaignAppProfilePageStateOn      CampaignAppProfilePageState = "ON"
	CampaignAppProfilePageStateOff     CampaignAppProfilePageState = "OFF"
)

// SpecialIndustry 特殊广告分类
type SpecialIndustry string

const (
	SpecialIndustryHousing    SpecialIndustry = "HOUSING"    // 房地产，房屋保险，抵押贷款或相关的广告
	SpecialIndustryEmployment SpecialIndustry = "EMPLOYMENT" // 工作机会，实习机会，职业认证项目或相关的广告
	SpecialIndustryCredit     SpecialIndustry = "CREDIT"     // 信用卡申请，汽车贷款，长期融资或相关的广告
)

// RFCampaignType 合约推广系列类型
type RFCampaignType string

const (
	RFCampaignTypeStandard RFCampaignType = "STANDARD" // 覆盖和频次推广系列
	RFCampaignTypePulse    RFCampaignType = "PULSE"    // TikTok Pulse 推广系列
	RFCampaignTypeTopView  RFCampaignType = "TOPVIEW"  // TopView 推广系列
)

// ObjectiveType 表示推广目标的枚举类型
type ObjectiveType string

const (
	ObjectiveAppPromotion   ObjectiveType = "APP_PROMOTION"
	ObjectiveWebConversions ObjectiveType = "WEB_CONVERSIONS"
	ObjectiveReach          ObjectiveType = "REACH"
	ObjectiveTraffic        ObjectiveType = "TRAFFIC"
	ObjectiveVideoViews     ObjectiveType = "VIDEO_VIEWS"
	ObjectiveProductSales   ObjectiveType = "PRODUCT_SALES"
	ObjectiveEngagement     ObjectiveType = "ENGAGEMENT"
	ObjectiveLeadGeneration ObjectiveType = "LEAD_GENERATION"
	ObjectiveRFReach        ObjectiveType = "RF_REACH"
	ObjectiveTopViewReach   ObjectiveType = "TOPVIEW_REACH"
)

// BuyingType 表示购买类型的枚举类型
type BuyingType string

const (
	BuyingTypeAuction            BuyingType = "AUCTION"
	BuyingTypeReservationRF      BuyingType = "RESERVATION_RF"
	BuyingTypeReservationTopView BuyingType = "RESERVATION_TOP_VIEW"
)

// Campaign 结构体包含广告系列的详细信息
type Campaign struct {
	AdvertiserID                string                      `json:"advertiser_id"`                   // 广告主 ID
	CampaignID                  string                      `json:"campaign_id"`                     // 推广系列 ID
	CampaignSystemOrigin        CampaignSystemOrigin        `json:"campaign_system_origin"`          // 推广系列来源
	IsSearchCampaign            bool                        `json:"is_search_campaign"`              // 是否为搜索推广系列
	IsSmartPerformanceCampaign  bool                        `json:"is_smart_performance_campaign"`   // 是否为自动化类型的推广系列
	CampaignName                string                      `json:"campaign_name"`                   // 推广系列名称
	CampaignType                CampaignType                `json:"campaign_type"`                   // 推广系列类型
	PostbackWindowMode          PostbackWindowMode          `json:"postback_window_mode"`            // SKAN 4.0 回传模式
	AppID                       string                      `json:"app_id"`                          // 推广的App的ID
	Budget                      float64                     `json:"budget"`                          // 推广系列预算
	BudgetMode                  BudgetMode                  `json:"budget_mode"`                     // 预算类型
	RtaID                       string                      `json:"rta_id"`                          // 实时 API ID
	RtaProductSelectionEnabled  bool                        `json:"rta_product_selection_enabled"`   // 是否使用实时 API 自动选择商品
	SecondaryStatus             string                      `json:"secondary_status"`                // 推广系列状态（二级状态）
	OperationStatus             OperationStatus             `json:"operation_status"`                // 推广系列的操作状态
	Objective                   Objective                   `json:"objective"`                       // 推广类型
	ObjectiveType               string                      `json:"objective_type"`                  // 推广目标
	AppPromotionType            AppPromotionType            `json:"app_promotion_type"`              // 应用推广类型
	CampaignProductSource       CampaignProductSource       `json:"campaign_product_source"`         // 推广系列的商品来源
	BudgetOptimizeOn            bool                        `json:"budget_optimize_on"`              // 是否开启推广系列预算优化
	BidType                     string                      `json:"bid_type"`                        // 推广系列层级的竞价策略
	DeepBidType                 string                      `json:"deep_bid_type"`                   // 深度事件出价类型
	RoasBid                     float64                     `json:"roas_bid,omitempty"`              // 用于价值优化的ROAS目标值
	OptimizationGoal            string                      `json:"optimization_goal"`               // 优化目标
	IsNewStructure              bool                        `json:"is_new_structure"`                // 是否为新结构
	CreateTime                  DateTime                    `json:"create_time"`                     // 推广系列创建时间
	ModifyTime                  DateTime                    `json:"modify_time"`                     // 推广系列修改时间
	IsAdvancedDedicatedCampaign bool                        `json:"is_advanced_dedicated_campaign"`  // 是否为高级专属推广系列
	CampaignAppProfilePageState CampaignAppProfilePageState `json:"campaign_app_profile_page_state"` // 下载中间页使用情况
	SpecialIndustries           *[]SpecialIndustry          `json:"special_industries"`              // 特殊广告分类
	RFCampaignType              RFCampaignType              `json:"rf_campaign_type"`                // 合约推广系列类型
}

// PageInfo 定义分页信息
type PageInfo struct {
	Page        int64 `json:"page"`         // 当前页数
	PageSize    int64 `json:"page_size"`    // 分页大小
	TotalNumber int64 `json:"total_number"` // 总结果数
	TotalPage   int64 `json:"total_page"`   // 总页数
}

// CampaignFilter 推广系列过滤条件
type CampaignFilter struct {
	ObjectiveType              ObjectiveType          `json:"objective_type,omitempty"`                // 推广目标，例如："APP_PROMOTION"
	CampaignIDs                []string               `json:"campaign_ids,omitempty"`                  // 推广系列 ID，允许数量范围: 1-100
	CampaignName               string                 `json:"campaign_name,omitempty"`                 // 推广系列名称，支持模糊匹配
	CampaignSystemOrigins      []CampaignSystemOrigin `json:"campaign_system_origins,omitempty"`       // 推广系列来源，枚举值：PROMOTE, TT_ADS_PLATFORM
	CampaignType               CampaignType           `json:"campaign_type,omitempty"`                 // 推广系列类型，枚举值: REGULAR_CAMPAIGN, IOS14_CAMPAIGN
	BuyingTypes                []BuyingType           `json:"buying_types,omitempty"`                  // 购买类型，枚举值：AUCTION, RESERVATION_RF, RESERVATION_TOP_VIEW
	PrimaryStatus              string                 `json:"primary_status,omitempty"`                // 一级状态
	SecondaryStatus            string                 `json:"secondary_status,omitempty"`              // 推广系列二级状态
	CreationFilterStartTime    string                 `json:"creation_filter_start_time,omitempty"`    // 推广系列最早创建时间，格式：YYYY-MM-DD HH:MM:SS（UTC时区）
	CreationFilterEndTime      string                 `json:"creation_filter_end_time,omitempty"`      // 推广系列最晚创建时间，格式：YYYY-MM-DD HH:MM:SS（UTC时区）
	IsSmartPerformanceCampaign bool                   `json:"is_smart_performance_campaign,omitempty"` // 是否为自动化类型的推广系列
	SplitTestEnabled           bool                   `json:"split_test_enabled,omitempty"`            // 推广系列是否启用了拆分对比测试
	CampaignProductSource      string                 `json:"campaign_product_source,omitempty"`       // 推广系列的商品来源，枚举值：CATALOG, STORE
	OptimizationGoal           string                 `json:"optimization_goal,omitempty"`             // 优化目标
}

// CampaignData 推广系列数据
type CampaignData struct {
	List     []Campaign `json:"list"`
	PageInfo PageInfo   `json:"page_info"`
}

// BaseReq 基础请求参数
type BaseReq struct {
	AdvertiserID                string   `json:"advertiser_id,omitempty"`
	Fields                      []string `json:"fields,omitempty"`
	ExcludeFieldTypesInResponse []string `json:"exclude_field_types_in_response,omitempty"`
	Page                        int64    `json:"page,omitempty" param:"default=1"`
	PageSize                    int64    `json:"page_size,omitempty" param:"default=10"`
}

// CampaignListReq 获取推广系列列表请求参数
type CampaignListReq struct {
	BaseReq
	Filtering CampaignFilter `json:"filtering,omitempty"`
}

// CampaignListResp 获取推广系列列表响应参数
type CampaignListResp struct {
	BaseResp
	Data CampaignData `json:"data"`
}

// FindCampaigns 获取推广系列列表 https://business-api.tiktok.com/portal/docs?id=1739665513181185
func (s *CampaignService) FindCampaigns(query *CampaignListReq) (*CampaignListResp, error) {
	apiUrl := "campaign/get/"
	resp := new(CampaignListResp)
	err := s.client.get(apiUrl, resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
