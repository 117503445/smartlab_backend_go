package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartlab/dto"
	"smartlab/model"
	"smartlab/serializer"

	"github.com/go-playground/validator/v10"
)

// DataLogCreate DataLog创建
func DataLogCreate(c *gin.Context) {
	dataLogIn := &dto.DataLogIn{}
	var err error
	if err = c.ShouldBindJSON(&dataLogIn); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(http.StatusBadRequest, "bad DataLogIn dto.", err))
		return
	}

	if err = validator.New().Struct(dataLogIn); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(serializer.StatusParamNotValid, "StatusParamNotValid", err))
		return
	}

	var dataLog *model.DataLog
	if dataLog, err = dataLogIn.ToDataLogIn(); err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Err(serializer.StatusDtoToModelError, "DTOtoModel failed", err))
		return
	}

	model.CreateDataLog(dataLog)
	c.JSON(http.StatusOK, dataLog)
}
