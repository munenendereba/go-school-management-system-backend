package models

import "time"

func (Grade) TableName() string {
	return "grade"
}

type Grade struct {
	Id        int64     `json:"id" gorm:"primarykey;autoincrement"`
	Grade     *string   `json:"grade" gorm:"column:grade"`
	NextGrade int64     `json:"nextGrade" gorm:"column:nextGrade"`
	CreatedAt time.Time `json:"createdAt"  gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt"  gorm:"column:updatedAt"`
}
