package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartlab/serializer"
	"smartlab/service"
)

func GetOpenID(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, serializer.Err(serializer.StatusParamNotValid, "code in query is nil", nil))
		return
	}
	if openid, err := service.WeChatAuthCode2Session(code); err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Err(serializer.StatusWeChatLoginError, "StatusWeChatLoginError", err) )
	} else {
		c.JSON(http.StatusOK, gin.H{"openid": openid})
	}
}
