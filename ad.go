package tba

type AdService service

type Ad struct {
	AdvertiserID               string                    `json:"advertiser_id,omitempty"`                  // 广告主 ID
	CampaignID                 string                    `json:"campaign_id,omitempty"`                    // 推广系列 ID
	CampaignName               string                    `json:"campaign_name,omitempty"`                  // 推广系列名称
	CampaignSystemOrigin       string                    `json:"campaign_system_origin,omitempty"`         // 推广系列来源
	AdgroupID                  string                    `json:"adgroup_id,omitempty"`                     // 广告组 ID
	AdgroupName                string                    `json:"adgroup_name,omitempty"`                   // 广告组名称
	AdID                       string                    `json:"ad_id,omitempty"`                          // 广告 ID
	AdName                     string                    `json:"ad_name,omitempty"`                        // 广告名称
	CreateTime                 string                    `json:"create_time,omitempty"`                    // 广告创建时间
	ModifyTime                 string                    `json:"modify_time,omitempty"`                    // 广告修改时间
	IdentityType               string                    `json:"identity_type,omitempty"`                  // 认证身份类型
	IdentityID                 string                    `json:"identity_id,omitempty"`                    // 认证身份ID
	IdentityAuthorizedBcID     string                    `json:"identity_authorized_bc_id,omitempty"`      // 商务中心ID
	CatalogID                  string                    `json:"catalog_id,omitempty"`                     // 商品库 ID @deprecated at 2024-06-30
	ProductSpecificType        string                    `json:"product_specific_type,omitempty"`          // 选择商品维度 枚举值：ALL,PRODUCT_SET,CUSTOMIZED_PRODUCTS
	ProductSetID               string                    `json:"product_set_id,omitempty"`                 // 商品系列 ID
	SkuIDs                     *[]string                 `json:"sku_ids,omitempty"`                        // SKU ID列表
	ShowcaseProducts           *[]ShowcaseProduct        `json:"showcase_products,omitempty"`              // 橱窗商品列表
	AdFormat                   string                    `json:"ad_format,omitempty"`                      // 广告样式 枚举值：SINGLE_IMAGE， SINGLE_VIDEO， LIVE_CONTENT, CAROUSEL_ADS（非视频购物类型的轮播广告），CATALOG_CAROUSEL（视频购物类型的轮播广告） 。
	VerticalVideoStrategy      string                    `json:"vertical_video_strategy,omitempty"`        // 视频类型 枚举值： UNSET（未设置），SINGLE_VIDEO（单个视频），CATALOG_VIDEOS（使用视频模板的商品库视频），CATALOG_UPLOADED_VIDEOS（使用已上传视频的商品库视频）， LIVE_STREAM（直播视频）。
	DynamicFormat              string                    `json:"dynamic_format,omitempty"`                 // 动态样式 枚举值：DYNAMIC_CREATIVE, UNSET
	VideoID                    string                    `json:"video_id,omitempty"`                       // 视频 ID
	ImageIDs                   *[]string                 `json:"image_ids,omitempty"`                      // 图片 ID列表
	CarouselImageIndex         int                       `json:"carousel_image_index,omitempty"`           // 轮播图片索引
	MusicID                    string                    `json:"music_id,omitempty"`                       // 音乐 ID
	TiktokItemID               string                    `json:"tiktok_item_id,omitempty"`                 // TikTok帖子 ID
	PromotionalMusicDisabled   bool                      `json:"promotional_music_disabled,omitempty"`     // 是否关闭推广音乐
	ItemDuetStatus             string                    `json:"item_duet_status,omitempty"`               // 合拍状态
	ItemStitchStatus           string                    `json:"item_stitch_status,omitempty"`             // 跟拍状态
	DarkPostStatus             string                    `json:"dark_post_status,omitempty"`               // dark post状态
	BrandedContentDisabled     bool                      `json:"branded_content_disabled,omitempty"`       // 品牌内容设置
	AdText                     string                    `json:"ad_text,omitempty"`                        // 广告文案
	AdTexts                    string                    `json:"ad_texts,omitempty"`                       // 广告文案列表
	CallToAction               string                    `json:"call_to_action,omitempty"`                 // 行动引导文案
	CallToActionID             string                    `json:"call_to_action_id,omitempty"`              // 行动引导文案素材包ID
	CardID                     string                    `json:"card_id,omitempty"`                        // 创意素材包 ID
	LandingPageURL             string                    `json:"landing_page_url,omitempty"`               // 落地页 URL
	UTMParams                  *[]UTMParam               `json:"utm_params,omitempty"`                     // URL参数列表
	PageID                     int                       `json:"page_id,omitempty"`                        // 页面ID
	CppURL                     string                    `json:"cpp_url,omitempty"`                        // 自定产品页面URL
	TikTokPageCategory         string                    `json:"tiktok_page_category,omitempty"`           // 您希望推广的 TikTok 页面的类型
	PhoneRegionCode            string                    `json:"phone_region_code,omitempty"`              //受众可点击广告拨打的电话号码的地区代码。
	PhoneRegionCallingCode     string                    `json:"phone_region_calling_code,omitempty"`      //受众可点击广告拨打的电话号码的区号。
	PhoneNumber                string                    `json:"phone_number,omitempty"`                   //受众可点击广告拨打的电话号码。
	Deeplink                   string                    `json:"deeplink,omitempty"`                       // 深度链接
	DeeplinkType               string                    `json:"deeplink_type,omitempty"`                  // 深度链接类型
	DeeplinkFormatType         string                    `json:"deeplink_format_type,omitempty"`           // 深度链接格式类型
	ShoppingAdsDeeplinkType    string                    `json:"shopping_ads_deeplink_type,omitempty"`     // 在购物广告中要使用的深度链接的来源。
	DeeplinkUtmParams          *[]UTMParam               `json:"deeplink_utm_params,omitempty"`            // 深度链接URL参数列表
	ShoppingAdsFallbackType    string                    `json:"shopping_ads_fallback_type,omitempty"`     // 在购物广告再营销场景下，深度链接唤起失败后的返回落地页类型。
	FallbackType               string                    `json:"fallback_type,omitempty"`                  // 失败路径。如果用户没有安装过应用，您可以选择让用户回退去安装应用，或者去您想推广的网页。
	DynamicDestination         string                    `json:"dynamic_destination,omitempty"`            // 动态目标页面URL,枚举值：DLP ,UNSET
	AigcDisclosureType         string                    `json:"aigc_disclosure_type,omitempty"`           // AIGC自主声明类型， SELF_DISCLOSURE，NOT_DECLARED
	DisclaimerType             string                    `json:"disclaimer_type,omitempty"`                // 免责声明类型，枚举值：TEXT_LINK（文字链免责声明），TEXT_ONLY（纯文本免责声明）。
	DisclaimerText             *DisclaimerText           `json:"disclaimer_text,omitempty"`                // 免责声明文本
	DisclaimerClickableTexts   *DisclaimerClickableTexts `json:"disclaimer_clickable_texts,omitempty"`     // 免责声明URL
	TrackingPixelID            int                       `json:"tracking_pixel_id,omitempty"`              // 正在监测的 Pixel ID
	TrackingAppID              string                    `json:"tracking_app_id,omitempty"`                // 应用ID
	TrackingOfflineEventSetIDs *[]string                 `json:"tracking_offline_event_set_ids,omitempty"` // 线下事件组ID列表
	TrackingMessageEventSetID  string                    `json:"tracking_message_event_set_id,omitempty"`  // 即时通讯广告中监测的消息事件集的 ID。
	VastMoatEnabled            bool                      `json:"vast_moat_enabled,omitempty"`              // 是否开启Moat（第三方视频可见性监测）。TikTok与Moat合作，为您在TikTok for Business上购买的品牌竞价、覆盖和频次信息流广告提供可见性监测。
	ViewabilityPostbidPartner  string                    `json:"viewability_postbid_partner,omitempty"`    // 出价后第三方可见性监测合作伙伴。枚举值：UNSET，MOAT，DOUBLE_VERIFY，IAS。
	ViewabilityVastURL         string                    `json:"viewability_vast_url,omitempty"`           // 出价后第三方监测合作伙伴用于监测可见性的封装的 VAST 网址。当brand_safety_postbid_partner非IAS时本字段为空值。
	BrandSafetyPostbidPartner  string                    `json:"brand_safety_postbid_partner,omitempty"`   // 出出价后第三方品牌安全监测合作伙伴。枚举值： UNSET，DOUBLE_VERIFY，IAS，ZEFR。
	BrandSafetyVastURL         string                    `json:"brand_safety_vast_url,omitempty"`          // 出价后第三方监测合作伙伴用于监测品牌安全的封装的 VAST 网址。当brand_safety_postbid_partner非IAS时本字段为空值。
	ImpressionTrackingURL      string                    `json:"impression_tracking_url,omitempty"`        // 展示监测URL
	ClickTrackingURL           string                    `json:"click_tracking_url,omitempty"`             // 点击监测URL
	PlayableUrl                string                    `json:"playable_url,omitempty"`                   // 试玩广告 URL。
	OperationStatus            string                    `json:"operation_status,omitempty"`               // 操作状态 枚举值：ENABLE，DISABLE，FROZEN
	SecondaryStatus            string                    `json:"secondary_status,omitempty"`               // 二级状态
	CreativeType               string                    `json:"creative_type,omitempty"`                  // 直播购物广告、商品购物广告或应用预注册场景中的创意素材类型
	AppName                    string                    `json:"app_name,omitempty"`                       // 应用名称
	DisplayName                string                    `json:"display_name,omitempty"`                   // 落地页或纯曝光广告的显示名称。
	AvatarIconWebUri           string                    `json:"avatar_icon_web_uri,omitempty"`            // 应用图标URL
	ProfileImageUrl            string                    `json:"profile_image_url,omitempty"`              // 头像URL
	CreativeAuthorized         bool                      `json:"creative_authorized,omitempty"`            // 是否允许在我们的 TikTok For Business 创意中心展示您的广告。仅对非美国广告主有效。默认值：false。
	IsAco                      bool                      `json:"is_aco,omitempty"`                         // 是否为程序化广告或智能创意广告。若为程序化广告或智能创意广告，则为true。
	IsNewStructure             bool                      `json:"is_new_structure,omitempty"`               // 是否为新结构
	OptimizationEvent          string                    `json:"optimization_event,omitempty"`             // 广告组转化事件
}

type DisclaimerText struct {
	Text string `json:"text,omitempty"` // 免责声明文本
}

type DisclaimerClickableTexts struct {
	Text string `json:"text,omitempty"` // 免责声明文本
	URL  string `json:"url,omitempty"`  // 免责声明URL
}

type ShowcaseProduct struct {
	ItemGroupID string `json:"item_group_id,omitempty"` // 商品SPU ID
	StoreID     string `json:"store_id,omitempty"`      // 店铺ID
	CatalogID   string `json:"catalog_id,omitempty"`    // 商品库ID
}

type UTMParam struct {
	Key   string `json:"key,omitempty"`   // 参数名
	Value string `json:"value,omitempty"` // 参数值
}

type AdFilter struct {
	CampaignIDs             []string `json:"campaign_ids,omitempty"`
	CampaignSystemOrigins   []string `json:"campaign_system_origins,omitempty"`
	AdgroupIDs              []string `json:"adgroup_ids,omitempty"`
	AdIDs                   []string `json:"ad_ids,omitempty"`
	PrimaryStatus           string   `json:"primary_status,omitempty"`
	SecondaryStatus         string   `json:"secondary_status,omitempty"`
	ObjectiveType           string   `json:"objective_type,omitempty"`
	BuyingTypes             []string `json:"buying_types,omitempty"`
	OptimizationGoal        string   `json:"optimization_goal,omitempty"`
	CreativeMaterialMode    string   `json:"creative_material_mode,omitempty"`
	Destination             string   `json:"destination,omitempty"`
	CreationFilterStartTime string   `json:"creation_filter_start_time,omitempty"`
	CreationFilterEndTime   string   `json:"creation_filter_end_time,omitempty"`
}

// AdListReq 定义了广告列表请求的参数结构
type AdListReq struct {
	BaseReq
	Filtering AdFilter `json:"filtering,omitempty"`
}

type CreateAdReq struct {
	AdvertiserID string `json:"advertiser_id,omitempty"` // 广告主 ID
	AdGroupID    string `json:"adgroup_id,omitempty"`    // 广告组 ID
	Creatives    []Ad   `json:"creatives,omitempty"`
}

type UpdateAdReq struct {
	AdvertiserID string `json:"advertiser_id,omitempty"` // 广告主 ID
	AdGroupID    string `json:"adgroup_id,omitempty"`    // 广告组 ID
	PatchUpdate  bool   `json:"patch_update,omitempty"`
	Creatives    []Ad   `json:"creatives,omitempty"`
}

// UpdateAdStatusReq 定义了更新广告状态请求的参数结构
type UpdateAdStatusReq struct {
	AdIDs           []string        `json:"ad_ids,omitempty"`           // 广告 ID 列表
	AcoAdIDs        []string        `json:"aco_ad_ids,omitempty"`       // ACO广告ID列表
	AdvertiserID    string          `json:"advertiser_id,omitempty"`    // 广告主 ID
	OperationStatus OperationStatus `json:"operation_status,omitempty"` // 要进行的操作状态
}

// AdListData 定义了广告数据结构
type AdListData struct {
	List     []Ad     `json:"list"`
	PageInfo PageInfo `json:"page_info"`
}

// AdListResp 定义了广告列表响应的结构
type AdListResp struct {
	BaseResp
	Data AdListData `json:"data"`
}

// AdData 定义了广告数据结构
type AdData struct {
	AdIDs     []string `json:"ad_ids"`
	Creatives []Ad     `json:"creatives"`
	PageInfo  PageInfo `json:"page_info"`
}

// AdResp 定义了广告响应的结构
type AdResp struct {
	BaseResp
	Data AdData `json:"data"`
}

// AdStatusData 定义了更新广告状态响应的数据结构
type AdStatusData struct {
	AdIDs    []string        `json:"ad_ids"`
	AcoAdIDs []string        `json:"aco_ad_ids"`
	Status   OperationStatus `json:"status"`
}

// AdStatusResp 定义了更新广告状态响应的结构
type AdStatusResp struct {
	BaseResp
	Data AdStatusData `json:"data"`
}

// FindAds 获取广告列表 https://business-api.tiktok.com/portal/docs?id=1735735588640770
func (s *AdService) FindAds(query *AdListReq) (*AdListResp, error) {
	apiUrl := "ad/get/"
	resp := new(AdListResp)
	err := s.client.get(apiUrl, resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateAd 创建广告 https://business-api.tiktok.com/portal/docs?id=1739953377508354
func (s *AdService) CreateAd(data *CreateAdReq) (*AdResp, error) {
	apiUrl := "ad/create/"
	resp := new(AdResp)
	err := s.client.post(apiUrl, resp, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateAd 更新广告 https://business-api.tiktok.com/portal/docs?id=1739953405142018
func (s *AdService) UpdateAd(data *UpdateAdReq) (*AdResp, error) {
	apiUrl := "ad/update/"
	resp := new(AdResp)
	err := s.client.post(apiUrl, resp, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateAdStatus 更新广告状态 https://business-api.tiktok.com/portal/docs?id=1739953422970882
func (s *AdService) UpdateAdStatus(req *UpdateAdStatusReq) (*AdStatusResp, error) {
	apiUrl := "ad/status/update/"
	resp := new(AdStatusResp)
	err := s.client.post(apiUrl, resp, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
