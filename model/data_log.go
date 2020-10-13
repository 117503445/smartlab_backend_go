package model

import "gorm.io/gorm"

// DataLog 实验数据日志
type DataLog struct {
	gorm.Model
	OpenID  string `json:"openid"`
	Page    string `json:"page"`
	Content string `json:"content"`
}
//CreateDataLog 保存DataLog
func CreateDataLog(datalog *DataLog) {
	DB.Save(datalog)
}
