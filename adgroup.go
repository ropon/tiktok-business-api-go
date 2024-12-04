package tba

type AdGroupService service

// AdGroup 定义了广告组的数据结构
type AdGroup struct {
	AdvertiserID                                 string             `json:"advertiser_id"`                                     // 广告主 ID
	CampaignID                                   string             `json:"campaign_id"`                                       // 推广系列 ID
	CampaignName                                 string             `json:"campaign_name"`                                     // 该广告组所属的推广系列的名称
	CampaignSystemOrigin                         string             `json:"campaign_system_origin"`                            // 广告组所属的推广系列来源
	AdGroupID                                    string             `json:"adgroup_id"`                                        // 广告组 ID
	AdGroupName                                  string             `json:"adgroup_name"`                                      // 广告组名称
	CreateTime                                   DateTime           `json:"create_time"`                                       // 广告组创建时间
	ModifyTime                                   DateTime           `json:"modify_time"`                                       // 广告组修改时间
	IsSmartPerformanceCampaign                   bool               `json:"is_smart_performance_campaign"`                     // 是否为自动化类型的推广系列中的广告组
	ShoppingAdsType                              string             `json:"shopping_ads_type"`                                 // 购物广告类型
	IdentityID                                   string             `json:"identity_id"`                                       // 认证身份 ID
	IdentityType                                 string             `json:"identity_type"`                                     // 认证身份类型
	IdentityAuthorizedBCID                       string             `json:"identity_authorized_bc_id"`                         // 认证身份绑定的商务中心的ID
	ProductSource                                string             `json:"product_source"`                                    // 商品来源
	CatalogID                                    string             `json:"catalog_id"`                                        // 商品库 ID
	CatalogAuthorizedBCID                        string             `json:"catalog_authorized_bc_id"`                          // 商务中心 ID
	StoreID                                      string             `json:"store_id"`                                          // TikTok Shop 店铺 ID
	StoreAuthorizedBCID                          string             `json:"store_authorized_bc_id"`                            // 商务中心的 ID
	PromotionType                                string             `json:"promotion_type"`                                    // 推广对象类型
	PromotionTargetType                          string             `json:"promotion_target_type"`                             // 推广目标线索收集的专属推广对象类型
	MessagingAppType                             string             `json:"messaging_app_type"`                                // 即时通讯应用类型
	MessagingAppAccountID                        string             `json:"messaging_app_account_id"`                          // 即时通讯应用账号 ID
	PhoneRegionCode                              []string           `json:"phone_region_code"`                                 // WhatsApp 电话号码的地区代码
	PhoneRegionCallingCode                       string             `json:"phone_region_calling_code"`                         // WhatsApp 电话号码的区号
	PhoneNumber                                  string             `json:"phone_number"`                                      // WhatsApp 电话号码
	PromotionWebsiteType                         string             `json:"promotion_website_type"`                            // TikTok 即时体验页面类型
	AppID                                        string             `json:"app_id"`                                            // 推广的App的ID
	AppType                                      string             `json:"app_type"`                                          // 推广的App的类型
	AppDownloadURL                               string             `json:"app_download_url"`                                  // App下载链接
	PixelID                                      string             `json:"pixel_id"`                                          // Pixel ID
	OptimizationEvent                            string             `json:"optimization_event"`                                // 广告组转化事件
	PlacementType                                string             `json:"placement_type"`                                    // 版位类型
	Placements                                   []string           `json:"placements"`                                        // 投放版位
	TikTokSubplacements                          []string           `json:"tiktok_subplacements"`                              // TikTok 子版位
	SearchResultEnabled                          bool               `json:"search_result_enabled"`                             // 是否在搜索广告中展示广告
	CommentDisabled                              bool               `json:"comment_disabled"`                                  // 是否允许用户评论广告
	VideoDownloadDisabled                        bool               `json:"video_download_disabled"`                           // 用户是否可以下载广告视频
	ShareDisabled                                bool               `json:"share_disabled"`                                    // 广告是否禁止分享到第三方平台
	BlockedPangleAppIDs                          []string           `json:"blocked_pangle_app_ids"`                            // Pangle 媒体屏蔽列表 ID
	AudienceType                                 string             `json:"audience_type"`                                     // 应用重定向受众类型
	AudienceRule                                 interface{}        `json:"audience_rule"`                                     // 受众规则
	AutoTargetingEnabled                         bool               `json:"auto_targeting_enabled"`                            // 是否开启自动定向
	ShoppingAdsRetargetingType                   string             `json:"shopping_ads_retargeting_type"`                     // 购物广告受众再营销类型
	ShoppingAdsRetargetingActionsDays            int                `json:"shopping_ads_retargeting_actions_days"`             // 受众行为的有效时间范围
	IncludedCustomActions                        []CustomAction     `json:"included_custom_actions"`                           // 包含的自定义行为
	ExcludedCustomActions                        []CustomAction     `json:"excluded_custom_actions"`                           // 排除的自定义行为
	ShoppingAdsRetargetingCustomAudienceRelation string             `json:"shopping_ads_retargeting_custom_audience_relation"` // 视频购物广告重定向受众与自定义受众间的逻辑关系
	LocationIDs                                  []string           `json:"location_ids"`                                      // 定向地域 ID
	ZipcodeIDs                                   []string           `json:"zipcode_ids"`                                       // 定向地域的邮政编码 ID
	Languages                                    []string           `json:"languages"`                                         // 受众语言
	Gender                                       string             `json:"gender"`                                            // 定向受众性别
	AgeGroups                                    []string           `json:"age_groups"`                                        // 定向受众年龄
	SpendingPower                                string             `json:"spending_power"`                                    // 消费能力
	HouseholdIncome                              []string           `json:"household_income"`                                  // 收入群体
	AudienceIDs                                  []string           `json:"audience_ids"`                                      // 受众 ID 列表
	SmartAudienceEnabled                         bool               `json:"smart_audience_enabled"`                            // 是否已启用智能受众
	ExcludedAudienceIDs                          []string           `json:"excluded_audience_ids"`                             // 排除受众 ID 列表
	InterestCategoryIDs                          []string           `json:"interest_category_ids"`                             // 一般兴趣分类 ID 列表
	InterestKeywordIDs                           []string           `json:"interest_keyword_ids"`                              // 其他兴趣分类 ID 列表
	PurchaseIntentionKeywordIDs                  []string           `json:"purchase_intention_keyword_ids"`                    // 购买意向分类 ID 列表
	Actions                                      []Action           `json:"actions"`                                           // 用于定向的行为分类对象数组
	SmartInterestBehaviorEnabled                 bool               `json:"smart_interest_behavior_enabled"`                   // 是否已启用智能型兴趣和行为
	IncludedPangleAudiencePackageIDs             []string           `json:"included_pangle_audience_package_ids"`              // 包含 Pangle 流量包 ID
	ExcludedPangleAudiencePackageIDs             []string           `json:"excluded_pangle_audience_package_ids"`              // 排除 Pangle 流量包 ID
	OperatingSystems                             []string           `json:"operating_systems"`                                 // 受众操作系统
	MinAndroidVersion                            string             `json:"min_android_version"`                               // 受众最低Android版本
	Ios14Targeting                               string             `json:"ios14_targeting"`                                   // iOS 设备版本
	MinIosVersion                                string             `json:"min_ios_version"`                                   // 受众最低 iOS 版本
	Ios14QuotaType                               string             `json:"ios14_quota_type"`                                  // iOS14专属推广系列配额
	DeviceModelIDs                               []string           `json:"device_model_ids"`                                  // 设备机型ID列表
	NetworkTypes                                 []string           `json:"network_types"`                                     // 网络类型
	CarrierIDs                                   []string           `json:"carrier_ids"`                                       // 运营商ID
	IspIDs                                       []string           `json:"isp_ids"`                                           // 互联网服务提供商的ID
	DevicePriceRanges                            []int              `json:"device_price_ranges"`                               // 设备价格区间
	TargetingExpansion                           TargetingExpansion `json:"targeting_expansion"`                               // 定向扩量设置
	SavedAudienceID                              string             `json:"saved_audience_id"`                                 // 已保存受众ID
	ContextualTagIDs                             []string           `json:"contextual_tag_ids"`                                // 内容相关定向标签ID列表
	BrandSafetyType                              string             `json:"brand_safety_type"`                                 // 品牌安全类型
	BrandSafetyPartner                           string             `json:"brand_safety_partner"`                              // 品牌安全合作伙伴
	InventoryFilterEnabled                       bool               `json:"inventory_filter_enabled"`                          // 库存筛选
	CategoryExclusionIDs                         []string           `json:"category_exclusion_ids"`                            // 内容排除类别ID
	VerticalSensitivityID                        string             `json:"vertical_sensitivity_id"`                           // 行业敏感内容控制类型ID
	BudgetMode                                   string             `json:"budget_mode"`                                       // 广告预算类型
	Budget                                       float64            `json:"budget"`                                            // 广告组预算
	ScheduledBudget                              float64            `json:"scheduled_budget"`                                  // 隔日预算值
	ScheduleType                                 string             `json:"schedule_type"`                                     // 广告投放时间类型
	ScheduleStartTime                            DateTime           `json:"schedule_start_time"`                               // 广告投放起始时间
	ScheduleEndTime                              DateTime           `json:"schedule_end_time"`                                 // 广告投放结束时间
	PredictImpression                            int                `json:"predict_impression"`                                // 预估展示量
	TopviewReachRange                            []int              `json:"topview_reach_range"`                               // 预估覆盖人数区间
	PreDiscountCPM                               float64            `json:"pre_discount_cpm"`                                  // 预估折前 CPM
	CPM                                          float64            `json:"cpm"`                                               // 预估折后 CPM
	DiscountType                                 string             `json:"discount_type"`                                     // 预算折扣类型
	DiscountAmount                               float64            `json:"discount_amount"`                                   // 预算折扣的固定金额
	DiscountPercentage                           float64            `json:"discount_percentage"`                               // 预算折扣的百分比
	PreDiscountBudget                            float64            `json:"pre_discount_budget"`                               // 预估折前预算
	ScheduleInfos                                []ScheduleInfo     `json:"schedule_infos"`                                    // 覆盖和频次广告组的广告投放信息
	Dayparting                                   string             `json:"dayparting"`                                        // 广告投放安排
	OptimizationGoal                             string             `json:"optimization_goal"`                                 // 优化目标
	SecondaryOptimizationEvent                   string             `json:"secondary_optimization_event"`                      // 次要优化目标
	MessageEventSetID                            string             `json:"message_event_set_id"`                              // 消息事件集的 ID
	Frequency                                    int                `json:"frequency"`                                         // 频次
	FrequencySchedule                            int                `json:"frequency_schedule"`                                // 频次天数
	BidType                                      string             `json:"bid_type"`                                          // 竞价策略
	BidPrice                                     float64            `json:"bid_price"`                                         // 出价
	ConversionBidPrice                           float64            `json:"conversion_bid_price"`                              // oCPM转化出价
	DeepBidType                                  string             `json:"deep_bid_type"`                                     // 深度事件出价类型
	RoasBid                                      float64            `json:"roas_bid"`                                          // ROAS 目标值
	VboWindow                                    string             `json:"vbo_window"`                                        // 竞价策略的时间窗口期
	BidDisplayMode                               string             `json:"bid_display_mode"`                                  // 每次浏览成本的计算和测算模式
	DeepCpaBid                                   float64            `json:"deep_cpa_bid"`                                      // 深度事件出价价格
	CpvVideoDuration                             string             `json:"cpv_video_duration"`                                // 视频播放时长
	NextDayRetention                             float64            `json:"next_day_retention"`                                // 次日留存率
	ClickAttributionWindow                       string             `json:"click_attribution_window"`                          // 点击归因窗口期
	EngagedViewAttributionWindow                 string             `json:"engaged_view_attribution_window"`                   // 深度互动观看归因窗口期
	ViewAttributionWindow                        string             `json:"view_attribution_window"`                           // 展示归因窗口期
	AttributionEventCount                        string             `json:"attribution_event_count"`                           // 事件统计方式
	BillingEvent                                 string             `json:"billing_event"`                                     // 计费事件
	Pacing                                       string             `json:"pacing"`                                            // 广告投放速度类型
	OperationStatus                              string             `json:"operation_status"`                                  // 广告组的操作状态
	SecondaryStatus                              string             `json:"secondary_status"`                                  // 广告组状态（二级状态)
	StatisticType                                string             `json:"statistic_type"`                                    // 转化事件出价统计类型
	IsHfss                                       bool               `json:"is_hfss"`                                           // 是否是 HFSS 食品
	CreativeMaterialMode                         string             `json:"creative_material_mode"`                            // 创意投放方式
	AdGroupAppProfilePageState                   string             `json:"adgroup_app_profile_page_state"`                    // 广告组是否使用了App中间页
	FeedType                                     string             `json:"feed_type"`                                         // 信息流类型
	RfPurchasedType                              string             `json:"rf_purchased_type"`                                 // 覆盖和频次广告购买方式
	PurchasedImpression                          int                `json:"purchased_impression"`                              // 覆盖和频次广告购买的展示量
	PurchasedReach                               int                `json:"purchased_reach"`                                   // 覆盖和频次广告购买的触达人数
	RfEstimatedCpr                               float64            `json:"rf_estimated_cpr"`                                  // 覆盖和频次广告预测的千人覆盖成本
	RfEstimatedFrequency                         float64            `json:"rf_estimated_frequency"`                            // 覆盖和频次广告预测的人均展示频次
	SplitTestGroupID                             string             `json:"split_test_group_id"`                               // 拆分对比测试组ID
	SplitTestStatus                              string             `json:"split_test_status"`                                 // 拆分对比测试状态
	IsNewStructure                               bool               `json:"is_new_structure"`                                  // 推广系列是否是新结构
	SkipLearningPhase                            bool               `json:"skip_learning_phase"`                               // 是否跳过学习阶段
}

// CustomAction 定义了自定义行为的数据结构
type CustomAction struct {
	Code string `json:"code"` // 自定义行为
	Days int    `json:"days"` // 时间区间
}

// Action 定义了用户行为的数据结构
type Action struct {
	ActionScene       string   `json:"action_scene"`        // 用户行为种类
	ActionPeriod      int      `json:"action_period"`       // 行为发生的时间
	VideoUserActions  []string `json:"video_user_actions"`  // 用户行为
	ActionCategoryIDs []string `json:"action_category_ids"` // 行为分类 ID 列表
}

// TargetingExpansion 定义了定向扩量的数据结构
type TargetingExpansion struct {
	ExpansionEnabled bool     `json:"expansion_enabled"` // 是否开启定向扩量
	ExpansionTypes   []string `json:"expansion_types"`   // 扩量的定向受众类型
}

// ScheduleInfo 定义了广告投放信息的结构
type ScheduleInfo struct {
	ScheduleID   string     `json:"schedule_id"`   // 广告投放信息ID
	IsDraft      bool       `json:"is_draft"`      // 广告投放信息是否为草稿
	Schedules    []Schedule `json:"schedules"`     // 广告的投放顺序
	DeliveryMode string     `json:"delivery_mode"` // 广告投放的排序与排期策略
}

// Schedule 定义了广告投放顺序和时间的结构
type Schedule struct {
	StartTime      DateTime `json:"start_time"`      // 广告投放开始时间
	EndTime        DateTime `json:"end_time"`        // 广告投放结束时间
	ExpectedOrders []int    `json:"expected_orders"` // 广告的投放顺序
}

// AdGroupData 定义了广告组数据结构
type AdGroupData struct {
	List     []AdGroup `json:"list,omitempty"`
	PageInfo PageInfo  `json:"page_info,omitempty"`
}

// AdGroupRequest 定义了广告组请求的参数结构
type AdGroupFilter struct {
	CampaignIDs             []string `form:"campaign_ids,omitempty"`               // 推广系列的id列表，支持筛选指定推广系列下的广告组，允许数量范围：1-100
	CampaignSystemOrigins   []string `form:"campaign_system_origins,omitempty"`    // 广告组所属的推广系列来源。枚举值：PROMOTE（内容加热推广系列），TT_ADS_PLATFORM（非内容加热推广系列）。默认值：["TT_ADS_PLATFORM"]
	AdGroupIDs              []string `form:"adgroup_ids,omitempty"`                // 广告组id列表，支持筛选指定的广告组，允许数量范围：1-100
	AdGroupName             string   `form:"adgroup_name,omitempty"`               // 广告组名字，支持广告组名字的模糊搜索
	PrimaryStatus           string   `form:"primary_status,omitempty"`             // 一级状态。枚举值详见枚举值 - 一级状态。默认值：STATUS_NOT_DELETE
	SecondaryStatus         string   `form:"secondary_status,omitempty"`           // 广告组二级状态。枚举值详见枚举值 - 二级状态 - 广告组状态
	ObjectiveType           string   `form:"objective_type,omitempty"`             // 推广目标。枚举值参见枚举值-推广目标
	BuyingTypes             []string `form:"buying_types,omitempty"`               // 购买类型。枚举值：AUCTION（竞价广告），RESERVATION_RF（合约覆盖和频次广告及TikTok Pulse广告），RESERVATION_TOP_VIEW（合约TopView广告）。默认值：["AUCTION", "RESERVATION_RF"]
	OptimizationGoal        string   `form:"optimization_goal,omitempty"`          // 优化目标。枚举值见枚举值-优化目标
	PromotionType           string   `form:"promotion_type,omitempty"`             // 推广对象类型（优化位置）。枚举值：APP（应用），WEBSITE（落地页），INSTANT_FORM（线索表单），LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE（TikTok私信），LEAD_GEN_CLICK_TO_SOCIAL_MEDIA_APP_MESSAGE（社交媒体应用），LEAD_GEN_CLICK_TO_CALL（电话通话）
	BidStrategy             string   `form:"bid_strategy,omitempty"`               // 竞价策略。枚举值：BID_STRATEGY_COST_CAP, BID_STRATEGY_BID_CAP, BID_STRATEGY_MAX_CONVERSION, BID_STRATEGY_LOWEST_COST
	CreativeMaterialMode    string   `form:"creative_material_mode,omitempty"`     // 创意投放方式。枚举值: CUSTOM（自定义），DYNAMIC（程序化），SMART_CREATIVE（智能创意）
	BillingEvents           []string `form:"billing_events,omitempty"`             // 计费事件，按照计费事件筛选。枚举值详见枚举值-计费事件
	CreationFilterStartTime string   `form:"creation_filter_start_time,omitempty"` // 广告组最早创建时间，格式：YYYY-MM-DD HH:MM:SS（UTC时区）
	CreationFilterEndTime   string   `form:"creation_filter_end_time,omitempty"`   // 广告组最晚创建时间，格式：YYYY-MM-DD HH:MM:SS（UTC时区）
	SplitTestEnabled        *bool    `form:"split_test_enabled,omitempty"`         // 是否启用了拆分对比测试筛选。true：仅获取已启用拆分对比测试的广告组，false：仅获取未启用拆分对比测试的广告组
}

// AdGroupListReq 定义了广告组列表请求的参数结构
type AdGroupListReq struct {
	BaseReq
	Filtering AdGroupFilter `json:"filtering"`
}

// AdGroupListResp 定义了广告组列表响应的结构
type AdGroupListResp struct {
	BaseResp
	Data AdGroupData `json:"data"`
}

// FindAdGroups 获取广告组列表 https://business-api.tiktok.com/portal/docs?id=1739314558673922
func (s *AdGroupService) FindAdGroups(query *AdGroupListReq) (*AdGroupListResp, error) {
	apiUrl := "adgroup/get/"
	resp := new(AdGroupListResp)
	err := s.client.get(apiUrl, resp, structPtr2Map(query, "form"))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
