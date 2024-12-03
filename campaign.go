package tba

type CampaignService service

type Campaign struct {
	AdvertiserID                string   `json:"advertiser_id"`
	CampaignID                  string   `json:"campaign_id"`
	CampaignSystemOrigin        *string  `json:"campaign_system_origin,omitempty"`
	IsSearchCampaign            bool     `json:"is_search_campaign"`
	IsSmartPerformanceCampaign  bool     `json:"is_smart_performance_campaign"`
	CampaignName                string   `json:"campaign_name"`
	CampaignType                string   `json:"campaign_type"`
	PostbackWindowMode          *string  `json:"postback_window_mode,omitempty"`
	AppID                       *string  `json:"app_id,omitempty"`
	Budget                      float64  `json:"budget"`
	BudgetMode                  string   `json:"budget_mode"`
	RtaID                       *string  `json:"rta_id,omitempty"`
	RtaProductSelectionEnabled  *bool    `json:"rta_product_selection_enabled,omitempty"`
	SecondaryStatus             string   `json:"secondary_status"`
	OperationStatus             string   `json:"operation_status"`
	Objective                   string   `json:"objective"`
	ObjectiveType               string   `json:"objective_type"`
	AppPromotionType            *string  `json:"app_promotion_type,omitempty"`
	CampaignProductSource       *string  `json:"campaign_product_source,omitempty"`
	BudgetOptimizeOn            *bool    `json:"budget_optimize_on,omitempty"`
	BidType                     *string  `json:"bid_type,omitempty"`
	DeepBidType                 *string  `json:"deep_bid_type,omitempty"`
	RoasBid                     *float64 `json:"roas_bid,omitempty"`
	OptimizationGoal            *string  `json:"optimization_goal,omitempty"`
	IsNewStructure              bool     `json:"is_new_structure"`
	CreateTime                  string   `json:"create_time"`
	ModifyTime                  string   `json:"modify_time"`
	IsAdvancedDedicatedCampaign bool     `json:"is_advanced_dedicated_campaign"`
	CampaignAppProfilePageState *string  `json:"campaign_app_profile_page_state,omitempty"`
	SpecialIndustries           []string `json:"special_industries,omitempty"`
	RfCampaignType              *string  `json:"rf_campaign_type,omitempty"`
}

type PageInfo struct {
	Page        int64 `json:"page"`
	PageSize    int64 `json:"page_size"`
	TotalNumber int64 `json:"total_number"`
	TotalPage   int64 `json:"total_page"`
}

type Filtering struct {
	CampaignIDs                []string `json:"campaign_ids"`
	CampaignName               string   `json:"campaign_name"`
	CampaignSystemOrigins      []string `json:"campaign_system_origins"`
	CampaignType               string   `json:"campaign_type"`
	ObjectiveType              string   `json:"objective_type"`
	BuyingTypes                []string `json:"buying_types"`
	PrimaryStatus              string   `json:"primary_status"`
	SecondaryStatus            string   `json:"secondary_status"`
	CreationFilterStartTime    string   `json:"creation_filter_start_time"`
	CreationFilterEndTime      string   `json:"creation_filter_end_time"`
	IsSmartPerformanceCampaign *bool    `json:"is_smart_performance_campaign"`
	SplitTestEnabled           *bool    `json:"split_test_enabled"`
	CampaignProductSource      string   `json:"campaign_product_source"`
	OptimizationGoal           string   `json:"optimization_goal"`
}

type CampaignData struct {
	List     []Campaign `json:"list,omitempty"`
	PageInfo PageInfo   `json:"page_info,omitempty"`
}

type CampaignListReq struct {
	AdvertiserID                string     `json:"advertiser_id"`
	Fields                      []string   `json:"fields"`
	ExcludeFieldTypesInResponse []string   `json:"exclude_field_types_in_response"`
	Filtering                   *Filtering `json:"filtering"`
	Page                        int64      `json:"page,default=1"`
	PageSize                    int64      `json:"page_size,default=10"`
}

type CampaignListResp struct {
	BaseResp
	CampaignData `json:"data"`
}

func (s *CampaignService) FindCampaigns(query *CampaignListReq) (*CampaignListResp, error) {
	apiUrl := "campaign/get/"
	resp := new(CampaignListResp)
	err := s.client.get(apiUrl, resp, structPtr2Map(query, "json"))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
