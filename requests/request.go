package requests

import (
	"github.com/go-pack/opentaobao"
	"net/url"
)

type TaobaoTbkScOrderGetRequest struct {
	opentaobao.Params

	//需返回的字段列表
	Fields string

	//订单查询开始时间
	StartTime string

	//订单查询时间范围，单位：秒，最小60，最大1200，如不填写，默认60。查询常规订单、三方订单、渠道，及会员订单时均需要设置此参数，直接通过设置pageSize,PageNo翻页查询数据即可
	Span string

	//第几页，默认1，1~100
	PageNo string

	//页大小，默认20，1~100
	PageSize string

	//订单状态，1:全部订单，3：订单结算，12：订单付款，13：订单失效，14：订单成功；订单查询类型为‘结算时间’时，只能查订单结算状态
	TkStatus string

	//订单查询类型，创建时间“createTime”，或结算时间“settleTime”。当查询渠道或会员运营订单时，建议入参创建时间“createTime”进行查询
	OrderQueryType string

	//订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1，通过设置订单场景类型，媒体可以查询指定场景下的订单信息，例如不设置，或者设置为1，表示查询常规订单，常规订单包含淘宝客所有的订单数据，含渠道，及会员运营订单，但不包含3方分成，及维权订单
	OrderScene string

	//订单数据统计类型，1:2方订单，2:3方订单，如果不设置，或者设置为1，表示2方订单
	OrderCountType string
}

func NewTaobaoTbkScOrderGetRequest() *TaobaoTbkScOrderGetRequest {
	return &TaobaoTbkScOrderGetRequest{Params: opentaobao.NewParams()}
}

func (request *TaobaoTbkScOrderGetRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.order.get"
}

func (request *TaobaoTbkScOrderGetRequest) GetApiParas() url.Values {
	params := url.Values{}
	for key, value := range request.GetParams() {
		params.Set(key, value)
	}
	return params
}

func (request *TaobaoTbkScOrderGetRequest) SetFields(Fields string) {
	request.Set("fields", Fields)
}

func (request *TaobaoTbkScOrderGetRequest) SetStartTime(StartTime string) {
	request.Set("start_time", StartTime)
}

func (request *TaobaoTbkScOrderGetRequest) SetSpan(Span string) {
	request.Set("span", Span)
}

func (request *TaobaoTbkScOrderGetRequest) SetPageNo(PageNo string) {
	request.Set("page_no", PageNo)
}

func (request *TaobaoTbkScOrderGetRequest) SetPageSize(PageSize string) {
	request.Set("page_size", PageSize)
}

func (request *TaobaoTbkScOrderGetRequest) SetTkStatus(TkStatus string) {
	request.Set("tk_status", TkStatus)
}

func (request *TaobaoTbkScOrderGetRequest) SetOrderQueryType(OrderQueryType string) {
	request.Set("order_query_type", OrderQueryType)
}

func (request *TaobaoTbkScOrderGetRequest) SetOrderScene(OrderScene string) {
	request.Set("order_scene", OrderScene)
}

func (request *TaobaoTbkScOrderGetRequest) SetOrderCountType(OrderCountType string) {
	request.Set("order_count_type", OrderCountType)
}
