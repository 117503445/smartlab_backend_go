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
