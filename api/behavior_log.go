package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"smartlab/dto"
	"smartlab/model"
	"smartlab/serializer"
)

// BehaviorLogCreate BehaviorLog创建
func BehaviorLogCreate(c *gin.Context) {
	behaviorLogIn := &dto.BehaviorLogIn{}
	var err error
	if err = c.ShouldBindJSON(&behaviorLogIn); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(http.StatusBadRequest, "bad DataLogIn dto.", err))
		return
	}

	if err = validator.New().Struct(behaviorLogIn); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(serializer.StatusParamNotValid, "StatusParamNotValid", err))
		return
	}

	var behaviorLog *model.BehaviorLog
	if behaviorLog, err = behaviorLogIn.ToBehaviorLog(); err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Err(serializer.StatusDtoToModelError, "DTOtoModel failed", err))
		return
	}

	model.CreateBehaviorLog(behaviorLog)
	c.JSON(http.StatusOK, behaviorLog)
}
