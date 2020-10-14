package model

import (
	"fmt"
	"gorm.io/gorm"
)

// DataLog 实验数据日志
type BehaviorLog struct {
	gorm.Model
	OpenID  string `json:"openid"`
	Page    string `json:"page"`
	Control string `json:"control"`
}

func GetBehaviorLogCSVHeader() string {
	return "OpenID,Page,Control,CreatedAt\n"
}

func (behaviorLog BehaviorLog) ToCSVLine() string {
	return fmt.Sprintf("%v,%v,%v,%v\n", behaviorLog.OpenID, behaviorLog.Page, behaviorLog.Control, behaviorLog.CreatedAt)
}

func BehaviorLogToCSV(behaviorLogs *[]BehaviorLog) string {
	csv := GetBehaviorLogCSVHeader()
	for _, behaviorLog := range *behaviorLogs {
		csv += behaviorLog.ToCSVLine()
	}
	return csv
}

//CreateDataLog 保存DataLog
func CreateBehaviorLog(behaviorLog *BehaviorLog) {
	DB.Save(behaviorLog)
}

func ReadAllBehaviorLog() *[]BehaviorLog {
	var behaviorLogs []BehaviorLog
	DB.Find(&behaviorLogs)
	return &behaviorLogs
}
