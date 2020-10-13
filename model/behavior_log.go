package model

import "gorm.io/gorm"

// DataLog 实验数据日志
type BehaviorLog struct {
	gorm.Model
	OpenID  string `json:"openid"`
	Page    string `json:"page"`
	Control string `json:"control"`
}

//CreateDataLog 保存DataLog
func CreateBehaviorLog(behaviorLog *BehaviorLog) {
	DB.Save(behaviorLog)
}
