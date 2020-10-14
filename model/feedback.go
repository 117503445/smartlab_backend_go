package model

import (
	"fmt"
	"gorm.io/gorm"
)

// DataLog 实验数据日志
type Feedback struct {
	gorm.Model
	OpenID       string `json:"openid"`
	Page         string `json:"page"`
	Content      string `json:"content"`
	ContactInfo  string `json:"contactInfo"`
	FeedbackType string `json:"type"`
}

//CreateFeedback 保存 feedback
func CreateFeedback(feedback *Feedback) {
	DB.Save(feedback)
}

func GetFeedbackCSVHeader() string {
	return "OpenID,Page,Content,ContactInfo,FeedbackType,CreatedAt\n"
}

func (feedback Feedback) ToCSVLine() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", feedback.OpenID, feedback.Page, feedback.Content, feedback.ContactInfo, feedback.FeedbackType, feedback.CreatedAt)
}

func FeedbackToCSV(feedbacks *[]Feedback) string {
	csv := GetFeedbackCSVHeader()
	for _, feedback := range *feedbacks {
		csv += feedback.ToCSVLine()
	}
	return csv
}

func ReadAllFeedback() *[]Feedback {
	var feedbacks []Feedback
	DB.Find(&feedbacks)
	return &feedbacks
}
