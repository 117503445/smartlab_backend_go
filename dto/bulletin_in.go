package dto

import (
	"github.com/devfeel/mapper"
	"smartlab/model"
)

type BulletinIn struct {
	ImageUrl  string `json:"imageUrl"`
	Title    string `json:"title"`
}

func (bulletinIn BulletinIn) ToBulletinIn() (*model.Bulletin, error) {
	bulletin := &model.Bulletin{}
	if err := mapper.AutoMapper(&bulletinIn, bulletin); err != nil {
		return bulletin, err
	}

	return bulletin, nil
}

