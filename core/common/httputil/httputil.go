package httputil

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"fabapi/core/common/json"
)

var logger = wlogging.MustGetLoggerWithoutName()

func PostJson(uri string, jsons interface{}) ([]byte, error) {

	jsonParam, errs := json.Marshal(jsons) //转换成JSON返回的是byte[]
	if errs != nil {
		logger.Error(errs.Error())
		return nil, errs
	}
	logger.Debugf("| %5s | %v  |\n%s", http.MethodPost, uri, jsonParam)
	now := time.Now()
	//发送请求
	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(string(jsonParam)))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	//响应
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Read failed:", err)
		return nil, err
	}
	logger.Debugf("| %5s | %v  | %v  |------over\n%s", http.MethodPost, uri, time.Now().Sub(now), string(response))
	//返回结果
	return response, nil
}

func PostForm(uri string, paras map[string][]string) ([]byte, error) {

	resp, err := http.PostForm(uri, url.Values(paras))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return body, nil

}

func Get(uri string) ([]byte, error) {
	logger.Debugf("| %5s | %v  |\n", http.MethodPost, uri)
	now := time.Now()
	resp, err := http.Get(uri)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Debugf("| %5s | %v  | %v  |------over\n%s", http.MethodPost, uri, time.Now().Sub(now), string(body))
	return body, nil

}

func EngineRequest(url string, data interface{}, ret interface{}) error {
	r, err := PostJson(url, data)
	if err != nil {
		logger.Error("引擎无法连接,服务错误", err)
		return err
	}

	err = json.Unmarshal(r, ret)
	if err != nil {
		logger.Error("引擎返回数据格式错误:", err)
		return err
	}
	return nil
}

func EngineGetRequest(url string, ret interface{}) error {
	r, err := Get(url)
	if err != nil {
		logger.Error("引擎无法连接,服务错误", err)
		return err
	}

	err = json.Unmarshal(r, ret)
	if err != nil {
		logger.Error("引擎返回数据格式错误:", err)
		return err
	}
	return nil
}
