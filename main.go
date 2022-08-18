package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"fioepq9.cn/checkin_ecycloud/config"
	"fioepq9.cn/checkin_ecycloud/logger"
	"fioepq9.cn/checkin_ecycloud/model"
	"github.com/roylee0704/gron"
	"github.com/sirupsen/logrus"
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


func Checkin(name, email, passwd string) {
	log := logger.L.WithField("name", name)
	client, err := Login(email, passwd)
	if err != nil {
		log.WithFields(logrus.Fields{
			"email": email,
			"passwd": passwd,
		}).WithError(err).Error("登录失败：网络问题 / 账号密码错误")
		return
	}
	resp, err := client.Post(
		config.C.Checkin.Url,
		"application/javascript; charset=utf-8",
		&bytes.Buffer{},
	)
	if err != nil {
		log.WithError(err).Error("http.Post failed, 签到失败：网络问题")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("ioutil.ReadAll failed, 签到失败：检查config.yaml")
		return
	}
	var checkinResp model.CheckinResponse
	err = json.Unmarshal(body, &checkinResp)
	if err != nil {
		log.WithFields(logrus.Fields{
			"response": string(body),
		}).WithError(err).Error("json.Unmarshal failed, 签到失败：检查config.yaml")
		return
	}
	log.Info(checkinResp.Msg)
}

type APP struct {}

func (app *APP) Do() {
	for _, u := range config.C.Users {
		Checkin(u.Name, u.Email, u.Passwd)
	}
	fmt.Println("===== ===== ===== Everything is Done ===== ===== =====")
}

func main() {
	app := &APP{}
	app.Do()

	cron := gron.New()
	cron.AddFunc(gron.Every(24 * time.Hour).At("06:00"), func() {
		app.Do()
	})
	cron.Start()
	<- make(chan struct{})
}
