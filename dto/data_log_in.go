package dto

import (
	"github.com/devfeel/mapper"
	"smartlab/model"
)

type DataLogIn struct {
	OpenID  string `json:"openid" validate:"required"`
	Page    string `json:"page" validate:"required"`
	Content string `json:"content" gorm:"size:1000"`
}

func (dataLogIn DataLogIn) ToDataLogIn() (*model.DataLog, error) {
	dataLog := &model.DataLog{}
	if err := mapper.AutoMapper(&dataLogIn, dataLog); err != nil {
		return dataLog, err
	}

	return dataLog, nil
}

