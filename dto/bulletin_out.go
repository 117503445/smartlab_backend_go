package dto

import (
	"github.com/devfeel/mapper"
	"smartlab/model"
)

type BulletinOut struct {
	ID       uint   `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Title    string `json:"title"`
}

func FromBulletin(bulletin *model.Bulletin) (*BulletinOut, error) {
	bulletinOut := &BulletinOut{}
	if err := mapper.AutoMapper(bulletin, bulletinOut); err != nil {
		return bulletinOut, err
	}
	bulletinOut.ID = bulletin.ID
	return bulletinOut, nil
}

func FromBulletins(bulletins []model.Bulletin) ([]BulletinOut, error) {
	bulletinOuts := make([]BulletinOut, len(bulletins))
	for i, bulletin := range bulletins {
		if bulletinOut, err := FromBulletin(&bulletin); err == nil {
			bulletinOuts[i] = *bulletinOut
		} else {
			return nil, err
		}
	}

	return bulletinOuts, nil
}
