package responses

import "encoding/json"

type MapData struct{


//商品Id
NumIid string `json:"num_iid"`



//商品淘客转链
ClickUrl string `json:"click_url"`


}
type TbkTpwdConvertResponseResults struct{

Data MapData `json:"map_data"`

}
type TbkTpwdConvertResponseBody struct {
Results TbkTpwdConvertResponseResults `json:"data"`
}
type TbkTpwdConvertResponse struct {
Body TbkTpwdConvertResponseBody `json:"tbk_tpwd_convert_response"`
}

func NewTbkTpwdConvertResponse(content []byte) *TbkTpwdConvertResponse {
response := &TbkTpwdConvertResponse{}
json.Unmarshal(content, response)
return response
}