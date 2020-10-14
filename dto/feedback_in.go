package dto

import (
	"github.com/devfeel/mapper"
	"smartlab/model"
)

type FeedbackIn struct {
	OpenID       string `json:"openid"`
	Page         string `json:"page"`
	Content      string `json:"content"`
	ContactInfo  string `json:"contactInfo"`
	FeedbackType string `json:"type"`
}

func (feedbackIn FeedbackIn) ToFeedback() (*model.Feedback, error) {
	feedback := &model.Feedback{}
	if err := mapper.AutoMapper(&feedbackIn, feedback); err != nil {
		return feedback, err
	}

	return feedback, nil
}
