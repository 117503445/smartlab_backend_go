package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"smartlab/conf"
	"smartlab/dto"
	"smartlab/model"
	"smartlab/router"
	"smartlab/serializer"
	"smartlab/util"
	"strings"
	"testing"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func InitCleanDatabase() {
	dbName := viper.GetString("mysql.dbname")
	if _, err := model.Exec(fmt.Sprintf("drop database %v", dbName)); err != nil {
		panic("删除数据库失败")
	}

	model.InitDatabase() //重新创建空白的数据库
}

func TestMain(m *testing.M) {
	conf.Init()
	dbName := viper.GetString("mysql.dbname")
	if !strings.Contains(dbName, "test") {
		panic("本测试会清空数据库,禁止在 数据库名 不包含 test 的 数据库上运行")
	}
	r = router.NewRouter()
	exitCode := m.Run()
	os.Exit(exitCode)
}
func TestPing(t *testing.T) {

	_, response := httpPost(t, r, "/api/ping", nil, "")

	expectResponse := "\"pong\""
	assert.Equal(t, expectResponse, response)
}

func TestUserRegister(t *testing.T) {

	userCreateUpdateIn := dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	code, response := httpPostJson(t, r, "/api/user", nil, userCreateUpdateIn)

	assert.Equal(t, http.StatusOK, code)

	expectResponse := gin.H{
		"id":       float64(2),
		"username": "user1",
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}

}
func TestUserRegisterParamNotValidError(t *testing.T) {

	userCreateIn := dto.UserCreateIn{
		UserName: "u",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	code, response := httpPostJson(t, r, "/api/user", nil, userCreateIn)

	assert.Equal(t, http.StatusBadRequest, code)

	expectResponse := gin.H{
		"code":    float64(serializer.StatusParamNotValid),
		"message": "StatusParamNotValid",
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}

}

func TestUserLogin(t *testing.T) {

	userCreateIn := dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userLoginDto := dto.UserLoginIn{
		UserName: "user1",
		Password: "pass1",
	}

	code, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)

	assert.Equal(t, http.StatusOK, code)

	expectResponse := gin.H{
		"code": float64(200),
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}

}

func TestUserMe(t *testing.T) {

	userCreateIn := dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userLoginDto := dto.UserLoginIn{
		UserName: "user1",
		Password: "pass1",
	}

	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)

	code, response := httpGetJson(t, r, "/api/user/me", map[string]string{"Authorization": authorization})
	assert.Equal(t, http.StatusOK, code)

	expectResponse := gin.H{
		"id":       float64(2),
		"username": "user1",
		"avatar":   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}

}
func TestUserMeUnauthorized(t *testing.T) {

	request, err := http.NewRequest(http.MethodGet, "/api/user/me", nil)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal([]byte(recorder.Body.String()), &response)
	assert.Nil(t, err)

	expectResponse := map[string]interface{}(gin.H{
		"code":    float64(401),
		"message": "cookie token is empty",
	})
	assert.Equal(t, expectResponse, response)
}

func TestCreateAdminPasswordTxt(t *testing.T) {
	filePath := util.FilePasswordAdmin
	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, 12, len(string(bytes)))
}
func TestCreateJwtPasswordTxt(t *testing.T) {
	filePath := util.FilePasswordJWT
	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, 12, len(string(bytes)))
}

func TestUserUpdate(t *testing.T) {

	userCreateIn := dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userLoginDto := dto.UserLoginIn{
		UserName: "user1",
		Password: "pass1",
	}

	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)

	userCreateIn = dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "newAva",
	}
	code, response := httpPutJson(t, r, "/api/user", map[string]string{"Authorization": authorization},
		userCreateIn)

	assert.Equal(t, http.StatusOK, code)
	expectResponse := gin.H{
		"id":       float64(2),
		"username": "user1",
		"avatar":   "newAva",
	}
	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}

}
func TestUserUpdateOptionalPassword(t *testing.T) {
	InitCleanDatabase()
	userCreateIn := dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userLoginIn := dto.UserLoginIn{
		UserName: "user1",
		Password: "pass1",
	}

	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginIn)
	authorization := "Bearer " + response["token"].(string)

	_, response = httpPutJson(t, r, "/api/user", map[string]string{"Authorization": authorization},
		gin.H{"password": "pass2"})

	userLoginIn = dto.UserLoginIn{
		UserName: "user1",
		Password: "pass2",
	}

	_, response = httpPostJson(t, r, "/api/user/login", nil, userLoginIn)

	expectResponse := gin.H{
		"code": float64(200),
	}
	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}

}
func TestUserUpdateRepeatUsernameError(t *testing.T) {
	InitCleanDatabase()
	userCreateIn := dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userCreateIn = dto.UserCreateIn{
		UserName: "user2",
		Password: "pass2",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userLoginDto := dto.UserLoginIn{
		UserName: "user1",
		Password: "pass1",
	}

	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)

	userCreateIn = dto.UserCreateIn{
		UserName: "user2",
		Password: "pass1",
		Avatar:   "newAva",
	}
	code, response := httpPutJson(t, r, "/api/user", map[string]string{"Authorization": authorization},
		userCreateIn)

	assert.Equal(t, http.StatusBadRequest, code)
	expectResponse := gin.H{
		"code":    float64(serializer.StatusUsernameRepeat),
		"message": "Username has already exists.",
	}
	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestUserUpdateParamNotValidError(t *testing.T) {
	InitCleanDatabase()
	userCreateIn := dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userCreateIn = dto.UserCreateIn{
		UserName: "user2",
		Password: "pass2",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userLoginDto := dto.UserLoginIn{
		UserName: "user1",
		Password: "pass1",
	}

	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)

	userCreateIn = dto.UserCreateIn{
		UserName: "u",
		Password: "pass1",
		Avatar:   "newAva",
	}
	code, response := httpPutJson(t, r, "/api/user", map[string]string{"Authorization": authorization},
		userCreateIn)

	assert.Equal(t, http.StatusBadRequest, code)
	expectResponse := gin.H{
		"code":    float64(serializer.StatusParamNotValid),
		"message": "username length should in [5,30]",
	}
	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}

func TestAdminUserRead(t *testing.T) {

	filePath := util.FilePasswordAdmin
	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	password := string(bytes)
	assert.Equal(t, 12, len(password))

	userLoginDto := dto.UserLoginIn{
		UserName: "admin",
		Password: password,
	}
	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)
	code, response := httpGetJson(t, r, "/api/admin/user/1", map[string]string{"Authorization": authorization})
	assert.Equal(t, http.StatusOK, code)

	expectResponse := gin.H{
		"id":       float64(1),
		"username": "admin",
		"avatar":   "",
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestAdminUserReadNotFoundError(t *testing.T) {
	InitCleanDatabase()

	filePath := util.FilePasswordAdmin
	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	password := string(bytes)
	assert.Equal(t, 12, len(password))

	userLoginDto := dto.UserLoginIn{
		UserName: "admin",
		Password: password,
	}
	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)
	code, response := httpGetJson(t, r, "/api/admin/user/2", map[string]string{"Authorization": authorization})
	assert.Equal(t, http.StatusNotFound, code)

	expectResponse := gin.H{
		"code":    float64(404),
		"message": "user not found",
		"error":   "record not found",
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestAdminUserReadNoRoleError(t *testing.T) {

	filePath := util.FilePasswordAdmin
	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	password := string(bytes)
	assert.Equal(t, 12, len(password))

	code, response := httpGetJson(t, r, "/api/admin/user/1", nil)
	assert.Equal(t, http.StatusUnauthorized, code)

	expectResponse := gin.H{
		"code":    float64(401),
		"message": "cookie token is empty",
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}

func TestDataLogCreate(t *testing.T) {
	var dataLogIn = dto.DataLogIn{
		OpenID:  "OpenID",
		Page:    "Page",
		Content: "Content",
	}
	code, response := httpPostJson(t, r, "/api/DataLog", nil, dataLogIn)
	assert.Equal(t, http.StatusOK, code)

	expectResponse := gin.H{
		"ID": float64(1),
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}

func TestBehaviorLogCreate(t *testing.T) {
	var behaviorLogIn = dto.BehaviorLogIn{
		OpenID:  "OpenID",
		Page:    "Page",
		Control: "Control",
	}
	code, response := httpPostJson(t, r, "/api/BehaviorLog", nil, behaviorLogIn)
	assert.Equal(t, http.StatusOK, code)

	expectResponse := gin.H{
		"ID": float64(1),
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestBehaviorLogViewCSV(t *testing.T) {
	var behaviorLogIn = dto.BehaviorLogIn{
		OpenID:  "OpenID",
		Page:    "Page",
		Control: "Control",
	}
	code, _ := httpPostJson(t, r, "/api/BehaviorLog", nil, behaviorLogIn)
	assert.Equal(t, http.StatusOK, code)

	code, responseText := httpGet(t, r, "/api/BehaviorLog/csv", nil)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "OpenID,Page,Control,CreatedAt\nOpenID,Page,Control", responseText[0:49])
}

func TestFeedbackCreate(t *testing.T) {
	var feedbackIn = dto.FeedbackIn{
		OpenID:       "OpenID",
		Page:         "Page",
		Content:      "Content",
		ContactInfo:  "ContactInfo",
		FeedbackType: "FeedbackType",
	}
	code, response := httpPostJson(t, r, "/api/feedback", nil, feedbackIn)
	assert.Equal(t, http.StatusOK, code)

	expectResponse := gin.H{
		"ID": float64(1),
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestFeedbackViewCSV(t *testing.T) {
	var feedbackIn = dto.FeedbackIn{
		OpenID:       "OpenID",
		Page:         "Page",
		Content:      "Content",
		ContactInfo:  "ContactInfo",
		FeedbackType: "FeedbackType",
	}
	code, _ := httpPostJson(t, r, "/api/feedback", nil, feedbackIn)
	assert.Equal(t, http.StatusOK, code)

	code, responseText := httpGet(t, r, "/api/feedback/csv", nil)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "OpenID,Page,Content,ContactInfo,FeedbackType,CreatedAt\nOpenID,Page,Content,ContactInfo,FeedbackType,", responseText[0:100])
}

func TestBulletinCreate(t *testing.T) {
	filePath := util.FilePasswordAdmin
	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	password := string(bytes)
	assert.Equal(t, 12, len(password))

	userLoginDto := dto.UserLoginIn{
		UserName: "admin",
		Password: password,
	}
	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)

	var bulletinIn = dto.BulletinIn{
		ImageUrl: "http://xd.117503445.top:8888/public/1.jpg",
		Title:    "hello",
	}
	code, response := httpPostJson(t, r, "/api/Bulletin",  map[string]string{"Authorization": authorization}, bulletinIn)
	assert.Equal(t, http.StatusOK, code)

	expectResponse := gin.H{
		"imageUrl": "http://xd.117503445.top:8888/public/1.jpg",
		"title":    "hello",
		"id":float64(1),
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestBulletinCreateNoRoleError(t *testing.T) {
	var bulletinIn = dto.BulletinIn{
		ImageUrl: "http://xd.117503445.top:8888/public/1.jpg",
		Title:    "hello",
	}
	code, response := httpPostJson(t, r, "/api/Bulletin", nil, bulletinIn)
	assert.Equal(t, http.StatusUnauthorized, code)

	expectResponse := gin.H{
		"message":    "cookie token is empty",
		"code":float64(401),
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestBulletinCreateUserNoRoleError(t *testing.T) {

	userCreateIn := dto.UserCreateIn{
		UserName: "user1",
		Password: "pass1",
		Avatar:   "https://gw.alicdn.com/tps/TB1W_X6OXXXXXcZXVXXXXXXXXXX-400-400.png",
	}

	httpPostJson(t, r, "/api/user", nil, userCreateIn)

	userLoginDto := dto.UserLoginIn{
		UserName: "user1",
		Password: "pass1",
	}

	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)

	var bulletinIn = dto.BulletinIn{
		ImageUrl: "http://xd.117503445.top:8888/public/1.jpg",
		Title:    "hello",
	}
	code, response := httpPostJson(t, r, "/api/Bulletin",  map[string]string{"Authorization": authorization}, bulletinIn)
	assert.Equal(t, http.StatusUnauthorized, code)

	expectResponse := gin.H{
		"message":    "has role failed: don't have role admin",
		"code":float64(401),
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestBulletinDeleteNoRoleError(t *testing.T) {
	filePath := util.FilePasswordAdmin
	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	password := string(bytes)
	assert.Equal(t, 12, len(password))

	userLoginDto := dto.UserLoginIn{
		UserName: "admin",
		Password: password,
	}
	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)

	var bulletinIn = dto.BulletinIn{
		ImageUrl: "http://xd.117503445.top:8888/public/1.jpg",
		Title:    "hello",
	}
	httpPostJson(t,r,"/api/Bulletin",   map[string]string{"Authorization": authorization}, bulletinIn )

	code, response := httpDeleteJson(t,r, "/api/Bulletin/1", nil, nil)
	assert.Equal(t, http.StatusUnauthorized, code)

	expectResponse := gin.H{
		"message":    "cookie token is empty",
		"code":float64(401),
	}

	for k := range expectResponse {
		assert.Equal(t, expectResponse[k], response[k])
	}
}
func TestBulletDelete(t *testing.T) {
	filePath := util.FilePasswordAdmin
	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	password := string(bytes)
	assert.Equal(t, 12, len(password))

	userLoginDto := dto.UserLoginIn{
		UserName: "admin",
		Password: password,
	}
	_, response := httpPostJson(t, r, "/api/user/login", nil, userLoginDto)
	authorization := "Bearer " + response["token"].(string)
	var bulletinIn = dto.BulletinIn{
		ImageUrl: "http://xd.117503445.top:8888/public/1.jpg",
		Title:    "hello",
	}
	if code, _ := httpPostJson(t,r,"/api/Bulletin",   map[string]string{"Authorization": authorization}, bulletinIn ); code == http.StatusOK{
		code, response := httpDeleteJson(t, r, "/api/Bulletin/1",  map[string]string{"Authorization": authorization},nil)
		assert.Equal(t, http.StatusOK, code)

		expectResponse := gin.H{
			"imageUrl": "http://xd.117503445.top:8888/public/1.jpg",
			"title":    "hello",
			"id":float64(1),
		}

		for k := range expectResponse {
			assert.Equal(t, expectResponse[k], response[k])
		}
	}

}
func httpRequest(t *testing.T, httpMethod string, router *gin.Engine, url string, headers map[string]string, body string) (responseCode int, responseText string) {
	request, err := http.NewRequest(httpMethod, url, strings.NewReader(body))
	assert.Nil(t, err)
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	return recorder.Code, recorder.Body.String()
}
func httpRequestJson(t *testing.T, httpMethod string, router *gin.Engine, url string, headers map[string]string, body interface{}) (responseCode int, responseMap map[string]interface{}) {
	js, _ := json.Marshal(body)
	code, text := httpRequest(t, httpMethod, router, url, headers, string(js))

	var response map[string]interface{}
	err := json.Unmarshal([]byte(text), &response)
	assert.Nil(t, err)

	return code, response
}

func httpGet(t *testing.T, router *gin.Engine, url string, headers map[string]string) (responseCode int, responseText string) {
	return httpRequest(t, http.MethodGet, router, url, headers, "")
}
func httpGetJson(t *testing.T, router *gin.Engine, url string, headers map[string]string) (responseCode int, responseMap map[string]interface{}) {
	return httpRequestJson(t, http.MethodGet, router, url, headers, "")
}

func httpPost(t *testing.T, router *gin.Engine, url string, headers map[string]string, body string) (responseCode int, responseText string) {
	return httpRequest(t, http.MethodPost, router, url, headers, body)
}
func httpPostJson(t *testing.T, router *gin.Engine, url string, headers map[string]string, body interface{}) (responseCode int, responseMap map[string]interface{}) {
	return httpRequestJson(t, http.MethodPost, router, url, headers, body)
}

func httpPut(t *testing.T, router *gin.Engine, url string, headers map[string]string, body string) (responseCode int, responseText string) {
	return httpRequest(t, http.MethodPut, router, url, headers, body)
}
func httpPutJson(t *testing.T, router *gin.Engine, url string, headers map[string]string, body interface{}) (responseCode int, responseMap map[string]interface{}) {
	return httpRequestJson(t, http.MethodPut, router, url, headers, body)
}

func httpDelete(t *testing.T, router *gin.Engine, url string, headers map[string]string, body string) (responseCode int, responseText string) {
	return httpRequest(t, http.MethodDelete, router, url, headers, body)
}
func httpDeleteJson(t *testing.T, router *gin.Engine, url string, headers map[string]string, body interface{}) (responseCode int, responseMap map[string]interface{}) {
	return httpRequestJson(t, http.MethodDelete, router, url, headers, body)
}
