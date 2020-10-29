package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"smartlab/dto"
	"smartlab/model"
	"smartlab/serializer"
	"strconv"
)

// BulletinCreate godoc
// @Summary BulletinCreate
// @Description 创建公告，需要管理员权限。
// @Accept  json
// @Produce  json
// @param BulletinIn body dto.BulletinIn true "dto.BulletinIn"
// @Success 200 {array} dto.BulletinOut
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
	bulletinOut, _ := dto.FromBulletin(bulletin)
	c.JSON(http.StatusOK, bulletinOut)
}

// BulletinReadAll godoc
// @Summary BulletinReadAll
// @Description 读取所有公告
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.BulletinOut
// @Router /Bulletin [get]
func BulletinReadAll(c *gin.Context) {
	bulletins := model.ReadAllBulletin()
	bulletinOut, _ := dto.FromBulletins(*bulletins)
	c.JSON(http.StatusOK, bulletinOut)
}
// BulletinDelete godoc
// @Summary BulletinDelete
// @Description 删除公告，需要管理员权限。
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.BulletinOut
// @param id query int true "DeleteBulletin.ID"
// @Security JWT
// @Router /Bulletin/{id} [delete]
func BulletinDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(http.StatusBadRequest, "id is not number", err))
		return
	}
	bulletin, err := model.ReadBulletinById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, serializer.Err(serializer.StatusDBError, "bulletin not found", err))
		return
	}

	if err := dto.DeleteBulletin(bulletin); err != nil{
		c.JSON(http.StatusInternalServerError, serializer.Err(serializer.StatusDBError, "bulletin not found", err))
		return
	}else{
		if BulletinOut, err := dto.FromBulletin(bulletin); err != nil {
			c.JSON(http.StatusInternalServerError, serializer.Err(serializer.StatusModelToDtoError, "dto.FromBulletin failed", err))
			return
		}else{
			c.JSON(http.StatusOK, BulletinOut)
			return
		}

	}

}