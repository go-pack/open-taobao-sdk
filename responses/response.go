package responses

import "encoding/json"



type NTbkOrder struct{


	//淘宝父订单号
	TradeParentId uint64 `json:"trade_parent_id"`



	//淘宝订单号
	TradeId uint64 `json:"trade_id"`



	//商品ID
	NumIid uint64 `json:"num_iid"`



	//商品标题
	ItemTitle string `json:"item_title"`



	//商品数量
	ItemNum uint64 `json:"item_num"`



	//单价
	Price string `json:"price"`



	//实际支付金额
	PayPrice string `json:"pay_price"`



	//卖家昵称
	SellerNick string `json:"seller_nick"`



	//卖家店铺名称
	SellerShopTitle string `json:"seller_shop_title"`



	//推广者获得的收入金额，对应联盟后台报表“预估收入”
	Commission string `json:"commission"`



	//推广者获得的分成比率，对应联盟后台报表“分成比率”
	CommissionRate string `json:"commission_rate"`



	//推广者unid（已废弃）
	Unid string `json:"unid"`



	//淘客订单创建时间
	CreateTime string `json:"create_time"`



	//淘客订单结算时间
	EarningTime string `json:"earning_time"`



	//淘客订单状态，3：订单结算，12：订单付款，13：订单失效，14：订单成功
	TkStatus uint64 `json:"tk_status"`



	//第三方服务来源，没有第三方服务，取值为“--”
	Tk3rdType string `json:"tk3rd_type"`



	//第三方推广者ID
	Tk3rdPubId uint64 `json:"tk3rd_pub_id"`



	//订单类型，如天猫，淘宝
	OrderType string `json:"order_type"`



	//收入比率，卖家设置佣金比率&#43;平台补贴比率
	IncomeRate string `json:"income_rate"`



	//效果预估，付款金额*(佣金比率&#43;补贴比率)*分成比率
	PubSharePreFee string `json:"pub_share_pre_fee"`



	//补贴比率
	SubsidyRate string `json:"subsidy_rate"`



	//补贴类型，天猫:1，聚划算:2，航旅:3，阿里云:4
	SubsidyType string `json:"subsidy_type"`



	//成交平台，PC:1，无线:2
	TerminalType string `json:"terminal_type"`



	//类目名称
	AuctionCategory string `json:"auction_category"`



	//来源媒体ID
	SiteId string `json:"site_id"`



	//来源媒体名称
	SiteName string `json:"site_name"`



	//广告位ID
	AdzoneId string `json:"adzone_id"`



	//广告位名称
	AdzoneName string `json:"adzone_name"`



	//付款金额
	AlipayTotalPrice string `json:"alipay_total_price"`



	//佣金比率
	TotalCommissionRate string `json:"total_commission_rate"`



	//佣金金额
	TotalCommissionFee string `json:"total_commission_fee"`



	//补贴金额
	SubsidyFee string `json:"subsidy_fee"`



	//渠道关系ID
	RelationId uint64 `json:"relation_id"`



	//会员运营id
	SpecialId uint64 `json:"special_id"`



	//跟踪时间
	ClickTime string `json:"click_time"`


}








type NTbkOrderResults struct{

	//淘宝客订单

	NTbkOrderList []NTbkOrder `json:"n_tbk_order"`


}




type TbkScOrderGetResponseBody struct {


	//淘宝客订单

	Results NTbkOrderResults `json:"results"`





}
type TbkScOrderGetResponse struct {
	Body TbkScOrderGetResponseBody `json:"tbk_sc_order_get_response"`
}

func NewTbkScOrderGetResponse(content []byte) *TbkScOrderGetResponse {
	response := &TbkScOrderGetResponse{}
	json.Unmarshal(content, response)
	return response
}