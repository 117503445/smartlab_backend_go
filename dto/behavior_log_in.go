package dto

import (
	"github.com/devfeel/mapper"
	"smartlab/model"
)

type BehaviorLogIn struct {
	OpenID  string `json:"openid" validate:"required"`
	Page    string `json:"page" validate:"required"`
	Control string `json:"control"`
}

func (behaviorLogIn BehaviorLogIn) ToBehaviorLog() (*model.BehaviorLog, error) {
	behaviorLog := &model.BehaviorLog{}
	if err := mapper.AutoMapper(&behaviorLogIn, behaviorLog); err != nil {
		return behaviorLog, err
	}

	return behaviorLog, nil
}
