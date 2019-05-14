package opentaobao_test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-pack/opentaobao"
	"github.com/go-pack/opentaobao/requests"
	"github.com/go-pack/opentaobao/responses"
	"github.com/polaris1119/goutils"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
	"testing"
)


func init()  {
	viper.SetConfigName(".config") //  设置配置文件名 (不带后缀)
	viper.AddConfigPath("/etc/test/")   // 第一个搜索路径
	viper.AddConfigPath("$HOME/.test")  // 可以多次调用添加路径
	viper.AddConfigPath(".")               // 比如添加当前目录
	viper.SetConfigType("yml")
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
func TestClient(t *testing.T) {

	session := viper.GetString("APP_TB_TOKEN")
	topClient := opentaobao.NewTopClient(viper.GetString("APP_TB_KEY"), viper.GetString("APP_TB_SECRET"))
	request := requests.NewTaobaoTbkScOrderGetRequest()
	request.SetFields("num_iid")
	//request.Set("fields", "tb_trade_parent_id,tb_trade_id,num_iid,item_title,item_num,price,pay_price,seller_nick,seller_shop_title,commission,commission_rate,unid,create_time,earning_time,tk3rd_pub_id,tk3rd_site_id,tk3rd_adzone_id,relation_id,tb_trade_parent_id,tb_trade_id,num_iid,item_title,item_num,price,pay_price,seller_nick,seller_shop_title,commission,commission_rate,unid,create_time,earning_time,tk3rd_pub_id,tk3rd_site_id,tk3rd_adzone_id,special_id,click_time")
	request.SetStartTime("2019-04-27 23:30:22")
	request.SetOrderQueryType("create_time")
	request.SetSpan("1200")
	content, err := topClient.Execute(request, session)
	if err != nil {
		println(err.Error())
		return
	}
	response := responses.NewTbkScOrderGetResponse(content)
	if response != nil && len(response.Body.Results.NTbkOrderList) > 0 {
		println(response.Body.Results.NTbkOrderList[0].NumIid)
	}
}
func TestClientData(t *testing.T) {
	session := viper.GetString("APP_TB_TOKEN")
	topClient := opentaobao.NewTopClient(viper.GetString("APP_TB_KEY"), viper.GetString("APP_TB_SECRET"))
	request := requests.NewTaobaoTbkTpwdConvertRequest()
	request.Set("siteId","390500317")//PID:mm_47170502_390500317_105849350469
	request.SetAdzoneId("105849350469")
	request.SetPasswordContent("iUH4YYCkbKN")
	content, err := topClient.Execute(request, session)
	if err != nil {
		println(err.Error())
		return
	}
	response := responses.NewTbkTpwdConvertResponse(content)
	if response != nil && len(response.Body.Results.Data.ClickUrl) > 0 {
		println(response.Body.Results.Data.ClickUrl)
	}
}

func TestAutoApi(t *testing.T) {
	res, err := http.Get("https://developer.alibaba.com/docs/api.htm?apiId=38078")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	doc.Find("#bd > div:nth-child(6) > table > tbody > tr.open-wrap2 > td > div > ul").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		text := s.Find("li.td-11").Text()
		name := strings.Split(text, "└")
		dataType := s.Find("li.td-12").Text()
		desc := s.Find("li.td-14").Text()
		camelName := goutils.CamelName(strings.Trim(name[1], "\r\n "))
		t.Logf("name %s type %s desc %s \r\n", camelName, strings.Trim(dataType, "\r\n "), strings.Trim(desc, "\r\n "))
	})
	println(doc.Find("#bd > div:nth-child(6) > table > tbody > tr.J_tableTr2.J_tableTrigger > td:nth-child(2)").Text())
}
func TestAutoApiResult(t *testing.T) {
	res, err := http.Get("https://developer.alibaba.com/docs/api.htm?spm=a219a.7395905.0.0.524575feo2JNYK&apiId=36836")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)

	text := doc.Find("#bd > div:nth-child(6) > table > tbody > tr.J_tableTr2.J_tableTrigger > td:nth-child(2)").Text()
	println(strings.Trim(strings.ReplaceAll(text," ",""), "\r\n "));
}

func TestGetResponseBodyName(t *testing.T) {

	req,_ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?apiId=38078")
	req.BuildRequest()

	println(req.GetResponseBodyName())

	if "tbk_sc_order_get_response" != req.GetResponseBodyName() {
		t.Error("req.GetResponseBodyName()")
	}
}
func TestGetResponseData(t *testing.T) {

	req,_ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?apiId=38078")
	req.BuildRequest()
	req.GetResponseData()
}
func TestGetResponseDataType(t *testing.T) {

	req,_ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?apiId=38078")
	req.BuildRequest()

}

func TestRequestList(t *testing.T) {

	req, _ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?apiId=38078")
	req.BuildRequest()
	req.Render(os.Stdout)
}


func TestResponseList(t *testing.T) {

	req,_ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?apiId=38078")
	req.BuildResponse()
	req.RenderResponse(os.Stdout)
}
func TestResponseData(t *testing.T) {
	req,_ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?apiId=38078")
	req.BuildResponse()
	req.RenderResponse(os.Stdout)
}
func TestMakeDataRequest(t *testing.T) {

	req,_ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?spm=a219a.7395905.0.0.6c7175feO9KvXv&apiId=32932")
	req.BuildRequest()
	req.Render(os.Stdout)
}

func TestRequestData(t *testing.T) {

	req, _ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?spm=a219a.7395905.0.0.5c8d75fe4XqZIo&apiId=28156")
	req.BuildRequest()
	req.Render(os.Stdout)
}
func TestMakeDataResponse(t *testing.T) {

	req,_ := opentaobao.NewMakeRequestInfo("https://developer.alibaba.com/docs/api.htm?spm=a219a.7395905.0.0.41c275feWodQur&apiId=31127")
	req.BuildResponse()
	req.RenderResponse(os.Stdout)
}