package model
import "gorm.io/gorm"

// Bulletin 公告数据
type Bulletin struct {
	gorm.Model
	ImageUrl  string `json:"imageUrl"`
	Title    string `json:"title"`

}

//CreateBulletin 保存Bulletin
func CreateBulletin(bulletin *Bulletin) {
	DB.Save(bulletin)
}

func ReadAllBulletin() *[]Bulletin {
	var bulletins []Bulletin
	DB.Find(&bulletins)
	return &bulletins
}

// ReadBulletinById 用ID获取用户
func ReadBulletinById(id int) (*Bulletin, error) {
	var bulletin Bulletin
	result := DB.First(&bulletin, id)
	return &bulletin, result.Error
}