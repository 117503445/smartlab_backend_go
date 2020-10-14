package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"smartlab/dto"
	"smartlab/model"
	"smartlab/serializer"
)

// FeedbackCreate Feedback 创建
func FeedbackCreate(c *gin.Context) {
	feedbackIn := &dto.FeedbackIn{}
	var err error
	if err = c.ShouldBindJSON(&feedbackIn); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(http.StatusBadRequest, "bad feedbackIn dto.", err))
		return
	}

	if err = validator.New().Struct(feedbackIn); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Err(serializer.StatusParamNotValid, "StatusParamNotValid", err))
		return
	}

	var feedback *model.Feedback
	if feedback, err = feedbackIn.ToFeedback(); err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Err(serializer.StatusDtoToModelError, "DTOtoModel failed", err))
		return
	}

	model.CreateFeedback(feedback)
	c.JSON(http.StatusOK, feedback)
}

func FeedbackViewCSV(c *gin.Context) {
	feedbacks := model.ReadAllFeedback()
	c.String(http.StatusOK, model.FeedbackToCSV(feedbacks))
}
