package opentaobao

import "net/url"

type Request interface {
	GetApiMethodName() string
	GetApiParas() url.Values
}

type Params map[string]string
func NewParams() Params {
	p := make(Params)
	return p
}
func (p Params) Set(key string, value string) {
	p[key] = value
}
func (p Params) GetParams() ( map[string]string) {
	return p
}
