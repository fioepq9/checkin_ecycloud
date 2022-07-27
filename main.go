package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"fioepq9.cn/checkin_ecycloud/config"
	"fioepq9.cn/checkin_ecycloud/logger"
)

func Login(email string, passwd string) (*http.Client, error) {
	var client http.Client
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.Jar = jar
	_, err = client.PostForm(
		config.C.Login.Url,
		url.Values{
			"code":   {""},
			"email":  {email},
			"passwd": {passwd},
		},
	)
	if err != nil {
		return nil, err
	}
	return &client, err
}

type checkinResponse struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

func checkin(email, passwd string, log *logger.Logger) {
	client, err := Login(email, passwd)
	if err != nil {
		log.Error("登录失败：网络问题 / 账号密码错误", err)
		panic(err)
	}
	resp, err := client.Post(
		config.C.Checkin.Url,
		"application/javascript; charset=utf-8",
		&bytes.Buffer{},
	)
	if err != nil {
		log.Error("签到失败：网络问题", err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("签到失败：检查config.yaml", err)
		panic(err)
	}
	var checkinResp checkinResponse
	err = json.Unmarshal(body, &checkinResp)
	if err != nil {
		log.Error("签到失败：检查config.yaml", err)
		panic(err)
	}
	log.Info(checkinResp.Msg)
}

func main() {
	for _, u := range config.C.Users {
		checkin(u.Email, u.Passwd, logger.NewLogger(
			fmt.Sprintf("./log/%s.txt", u.Shortcut),
		))
	}
}
