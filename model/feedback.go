package model

import "gorm.io/gorm"

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
