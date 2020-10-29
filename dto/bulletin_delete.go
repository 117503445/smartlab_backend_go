package dto

import "smartlab/model"

func DeleteBulletin (bulletin *model.Bulletin )error{
	result := model.DB.Delete(bulletin)
	return result.Error
}
