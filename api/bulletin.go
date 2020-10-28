package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"smartlab/dto"
	"smartlab/model"
	"smartlab/serializer"
)

// BulletinCreate godoc
// @Summary BulletinCreate
// @Description 创建公告，需要管理员权限。
// @Accept  json
// @Produce  json
// @param BulletinIn body dto.BulletinIn true "dto.BulletinIn"
// @Success 200 {array} model.Bulletin
// @Security JWT
// @Router /Bulletin [post]
func BulletinCreate(c *gin.Context) {
	bulletinIn := &dto.BulletinIn{}
	var err error
	if err = c.ShouldBindJSON(&bulletinIn); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(http.StatusBadRequest, "bad bulletinIn dto.", err))
		return
	}

	if err = validator.New().Struct(bulletinIn); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(serializer.StatusParamNotValid, "StatusParamNotValid", err))
		return
	}

	var bulletin *model.Bulletin
	if bulletin, err = bulletinIn.ToBulletinIn(); err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Err(serializer.StatusDtoToModelError, "DTOtoModel failed", err))
		return
	}

	model.CreateBulletin(bulletin)
	c.JSON(http.StatusOK, bulletin)
}
// BulletinReadAll godoc
// @Summary BulletinReadAll
// @Description 读取所有公告
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Bulletin
// @Router /Bulletin [get]
func BulletinReadAll(c *gin.Context) {
	bulletins := model.ReadAllBulletin()
	c.JSON(http.StatusOK, bulletins)
}
