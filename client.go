package opentaobao

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func NewTopClient(AppKey string, SecretKey string) *TopClient  {
	return &TopClient{
		AppKey:AppKey,
		SecretKey:SecretKey,
		GatewayUrl:"http://gw.api.taobao.com/router/rest",
		ApiVersion:"2.0",
		Format:"json",
		SignMethod:"md5",
		ConnectTimeout: 5,

	}
}
type TopClient struct {
	AppKey         string
	SecretKey      string
	GatewayUrl     string
	Format         string
	SignMethod     string
	ApiVersion     string
	Session     string
	ConnectTimeout int
	ReadTimeout    int
}

func (client *TopClient) SetSession(session string)  {
	client.Session = session
}
func (client *TopClient) Execute(request Request, session string) (content []byte,err error) {
	requestParams := url.Values{}
	requestParams.Set("method", request.GetApiMethodName())
	if len(session) > 0 {
		requestParams.Set("session", session)
	}
	apiParams := request.GetApiParas()
	for key, value := range apiParams {
		requestParams.Set(key, value[0])
	}
	params := client.createQueryParams(&requestParams)

	req, err := http.NewRequest("POST", client.GatewayUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	proxy, _ := url.Parse("http://127.0.0.1:8888")
	netTransport := &http.Transport{Proxy:http.ProxyURL(proxy)}

	httpClient := &http.Client{Transport:netTransport}
	httpClient.Timeout = time.Duration(client.ConnectTimeout)*time.Second

	response, err := httpClient.Do(req)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	res, err := simplejson.NewJson(body)
	if err != nil {
		return
	}
	if responseError, ok := res.CheckGet("error_response"); ok {
		errorBytes, _ := responseError.Encode()
		err = errors.New("执行错误:" + string(errorBytes))
	}
	return body, err
}
func (client *TopClient) createQueryParams(p *url.Values) (url.Values) {
	// 公共参数
	args := url.Values{}
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	args.Add("timestamp", strconv.FormatInt(loc.Unix(), 10))
	args.Add("format", client.Format)
	args.Add("app_key", client.AppKey)
	args.Add("v", client.ApiVersion)
	args.Add("sign_method", client.SignMethod)
	args.Add("partner_id", "Undesoft")
	// 请求参数
	for key, val := range *p {
		args.Set(key, val[0])
	}
	// 设置签名
	args.Add("sign", client.generateSign(args))
	return args
}

func (client *TopClient) generateSign(args url.Values) string {
	// 获取Key
	keys := []string{}
	for k := range args {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	query := client.SecretKey
	for _, k := range keys {
		query += k + args.Get(k)
	}
	query += client.SecretKey
	// 使用MD5加密
	signBytes := md5.Sum([]byte(query))
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(signBytes[:]))
}
