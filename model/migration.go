package model

import "smartlab/util"

//执行数据迁移

func migration() {
	// 自动迁移模式
	if err := DB.AutoMigrate(&User{}); err != nil {
		util.Log().Error("AutoMigrate User Failed", err)
	}
	if err := DB.AutoMigrate(&DataLog{}); err != nil {
		util.Log().Error("AutoMigrate DataLog Failed", err)
	}
	if err := DB.AutoMigrate(&BehaviorLog{}); err != nil {
		util.Log().Error("AutoMigrate BehaviorLog Failed", err)
	}
	if err := DB.AutoMigrate(&Feedback{}); err != nil {
		util.Log().Error("AutoMigrate Feedback Failed", err)
	}
	if err := DB.AutoMigrate(&Bulletin{}); err != nil {
		util.Log().Error("AutoMigrate Bulletin Failed", err)
	}
}
