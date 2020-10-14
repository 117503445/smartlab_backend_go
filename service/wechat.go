package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/spf13/viper"
)

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func WeChatAuthCode2Session(code string) (string, error) {
	appid := viper.GetString("wechat.appid")
	secret := viper.GetString("wechat.secret")
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", appid, secret, code)
	resp, err := requests.Get(url)
	if err != nil {
		return "", err
	}
	text := resp.Text()
	var response map[string]interface{}
	err = json.Unmarshal([]byte(text), &response)
	if err != nil {
		return "", err
	}
	openid := response["openid"]
	if openid == nil {
		return "", errors.New("openid not found in response")
	} else {
		return openid.(string), nil
	}
}
