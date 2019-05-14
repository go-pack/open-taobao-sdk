package opentaobao

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/polaris1119/goutils"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

type RequestDataDesc struct {
	Type     string
	Label    string
	Desc     string
	ParamKey string
	Setter           []RequestDataDesc
}
type RespDataDesc struct {
	Type     string
	PublicType     string
	Label    string
	Desc     string
	JsonName string
	DataStructType string
	DataStructJsonName string
	IsBaseType bool
	CustomStructJsonName string
	Next *RespDataDesc
	Setter           []RespDataDesc
}
type MakeRequestInfo struct {
	Url              string
	Name             string
	ResultType       string
	doc              *goquery.Document
	Setter           []RequestDataDesc
	ResponseBodyName string
	ResponseData           []RespDataDesc
	CustomType map[string]*RespDataDesc
}

func NewMakeRequestInfo(url string) (*MakeRequestInfo, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, err
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	return &MakeRequestInfo{Url: url, doc: doc,CustomType:make(map[string]*RespDataDesc)}, nil
}

func (info *MakeRequestInfo) BuildRequest() (err error) {
	res, err := http.Get(info.Url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return err
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	info.Setter = make([]RequestDataDesc, 0)
	doc.Find("#bd > div:nth-child(4) > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.Find("td:nth-child(1) > span").Remove()
		Label := CamelName(s.Find("td:nth-child(1)").Text())
		//dataType := TopTypeToGo(CamelName(s.Find("td:nth-child(2)").Text()))
		dataType := "string"
		Desc := CamelName(s.Find("td:nth-child(6)").Text())
		info.Setter = append(info.Setter, RequestDataDesc{Label: Label, Type: dataType, Desc: Desc, ParamKey: goutils.UnderscoreName(Label)})
	})
	doc.Find("#pages > div > div.wrap-inner.block-docs-wrap.J_FloatContainer > div.docs-right > div.mtl > h2 > span").Remove()
	info.Name = doc.Find("#pages > div > div.wrap-inner.block-docs-wrap.J_FloatContainer > div.docs-right > div.mtl > h2").Text()

	return nil
}

func (info *MakeRequestInfo) SetName(name string) {
	info.Name = name
}
func (info *MakeRequestInfo) SetParam(Type string, Label string) {
	info.Setter = append(info.Setter, RequestDataDesc{Type: Type, Label: Label})
}
func (info *MakeRequestInfo) Render(wr io.Writer) {

	dir, _ := os.Getwd()
	fileName := dir + "/template/request.tmpl"
	tmpl, _ := template.ParseFiles(fileName)
	data := make(map[string]interface{})

	data["setters"] = info.Setter
	data["ApiNameRaw"] = info.Name
	data["ApiNameCamelNaming"] = CamelApiName(info.Name)
	tmpl.Execute(os.Stdout, data)
}
func (resp *MakeRequestInfo) RenderResponse(wr io.Writer) {

	dir, _ := os.Getwd()
	fileName := dir + "/template/response.tmpl"
	tmpl, _ := template.ParseFiles(fileName)
	data := make(map[string]interface{})

	data["CustomType"] = resp.CustomType
	data["ResponseData"] = resp.GetResponseData()
	data["ApiNameRaw"] = resp.ResponseBodyName
	data["ApiNameCamelNaming"] = CamelApiName(resp.ResponseBodyName)
	tmpl.Execute(os.Stdout, data)
}
func (resp *MakeRequestInfo) GetResponseBodyName() string {
	resp.ResponseBodyName = strings.TrimPrefix(resp.Name, "taobao.")
	all := strings.ReplaceAll(resp.ResponseBodyName, ".", "_")
	return strings.Trim(all, "\r\n ") + "_response";
}



func (resp *MakeRequestInfo) parseResultStruct(typeName string)  string {
	if strings.HasSuffix(typeName, "[]") {
		return "list"
	}else if IsBaseType(typeName){
		return "base"
	}else {
		return "object"
	}
}

func (resp *MakeRequestInfo)getCustomType(typeDesc string) string  {
	return strings.TrimSuffix(typeDesc,"[]")
}
func (resp *MakeRequestInfo) GetResponseData()  []RespDataDesc {
	resp.doc.Find("#bd > div:nth-child(6) > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("class")
		if val == "open-wrap2" {
			return
		}
		// For each item found, get the band and title
		s.Find("td:nth-child(1) > span").Remove()
		topName := TrimChar(s.Find("td:nth-child(1)").Text())
		topType := TrimChar(s.Find("td:nth-child(2)").Text())
		var dataType string
		isBaseType := IsBaseType(topType)
		if isBaseType {
			dataType = "string"
		}else{
			dataType = resp.getCustomType(topType)
		}

		Desc := TrimChar(s.Find("td:nth-child(4)").Text())
		resultStruct := resp.parseResultStruct(topType)

		dataTypeDesc := RespDataDesc{DataStructJsonName: goutils.UnderscoreName(topName),CustomStructJsonName:goutils.UnderscoreName(dataType),JsonName: goutils.UnderscoreName(topName),IsBaseType: isBaseType,PublicType:CamelName(dataType), Label: CamelName(topName), Type: dataType, Desc: Desc, DataStructType: resultStruct}
		if resultStruct == "list" || resultStruct == "object"{
			resp.ParseChildrenDataType(&dataTypeDesc,s.Siblings())
		}
		resp.ResponseData = append(resp.ResponseData, dataTypeDesc)

	})
	return resp.ResponseData
}
func (resp *MakeRequestInfo) ParseChildrenDataType(data *RespDataDesc,selectNode *goquery.Selection)  {
	selectNode.Find("ul").Each(func(i int, s *goquery.Selection) {
		text := TrimChar(s.Find("li.td-11").Text())
		name := strings.Split(text, "└")
		topType := TrimChar(s.Find("li.td-12").Text())
		Desc := TrimChar(s.Find("li.td-14").Text())
		topName := TrimChar(name[1])
		camelName := goutils.CamelName(topName)

		isBaseType := IsBaseType(topType)
		resultStruct := resp.parseResultStruct(topType)
		if isBaseType {
			topType = TopTypeToGo(topType)
		}
		dataTypeDesc := RespDataDesc{DataStructJsonName: goutils.UnderscoreName(topName),JsonName: goutils.UnderscoreName(topName),IsBaseType: isBaseType, Label: camelName, Type: topType, Desc: Desc, DataStructType: resultStruct}
		if resultStruct == "list" || resultStruct == "object"{
			resp.ParseChildrenDataType(&dataTypeDesc,s.Siblings())
		}
		data.Setter = append(data.Setter, dataTypeDesc)
	})
	resp.CustomType[data.Type] = data

}

func (resp *MakeRequestInfo) BuildResponse() {

	resp.doc.Find("#pages > div > div.wrap-inner.block-docs-wrap.J_FloatContainer > div.docs-right > div.mtl > h2 > span").Remove()
	resp.Name = resp.doc.Find("#pages > div > div.wrap-inner.block-docs-wrap.J_FloatContainer > div.docs-right > div.mtl > h2").Text()

	resp.ResponseBodyName = resp.GetResponseBodyName()

}


func (resp *MakeRequestInfo) parseCustomType()  {

}


func TrimChar(name string) string  {
	return strings.ReplaceAll(strings.Trim(name,"\r\n "), " ","")
}

func IsBaseType(topType string)  bool{
	if topType == "String" {
		return true
	}
	if topType == "Number" {
		return true
	}
	if topType == "Date" {
		return true
	}
	return false
}
func TopTypeToGo(topType string) (goType string) {
	if topType == "String" {
		return "string"
	}
	if topType == "Number" {
		return "uint64"
	}
	if topType == "Date" {
		return "string"
	}
	panic("topType " + topType + " not support!")
}

func CamelName(name string) string {
	name = strings.Trim(name, "\r\n \t");
	return goutils.CamelName(name)
}
func CamelApiName(name string) string {
	return goutils.CamelName(strings.ReplaceAll(name, ".", "_"))
}
