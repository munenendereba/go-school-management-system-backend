package models

func (Grade) TableName() string {
	return "grade"
}

type Grade struct {
	Id    int64   `json:"id" gorm:"primarykey;autoincrement"`
	Grade *string `json:"grade" gorm:"column:grade"`
}
