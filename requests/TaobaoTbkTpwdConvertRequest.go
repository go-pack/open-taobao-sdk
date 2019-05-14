package requests

import (
	"github.com/go-pack/opentaobao"
	"net/url"
)

type TaobaoTbkTpwdConvertRequest struct {
	opentaobao.Params

	//需要解析的淘口令
	PasswordContent string

	//广告位ID
	AdzoneId string

	//1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接
	Dx string
}

func NewTaobaoTbkTpwdConvertRequest() *TaobaoTbkTpwdConvertRequest {
	return &TaobaoTbkTpwdConvertRequest{Params: opentaobao.NewParams()}
}

func (request *TaobaoTbkTpwdConvertRequest) GetApiMethodName() string {
	return "taobao.tbk.tpwd.convert "
}

func (request *TaobaoTbkTpwdConvertRequest) GetApiParas() url.Values {
	params := url.Values{}
	for key, value := range request.GetParams() {
		params.Set(key, value)
	}
	return params
}

func (request *TaobaoTbkTpwdConvertRequest) SetPasswordContent(PasswordContent string) {
	request.Set("password_content", PasswordContent)
}

func (request *TaobaoTbkTpwdConvertRequest) SetAdzoneId(AdzoneId string) {
	request.Set("adzone_id", AdzoneId)
}

func (request *TaobaoTbkTpwdConvertRequest) SetDx(Dx string) {
	request.Set("dx", Dx)
}
