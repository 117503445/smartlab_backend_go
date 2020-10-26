package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"smartlab/dto"
	"smartlab/model"
	"smartlab/serializer"
)

// BulletinCreate Bulletin创建
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
