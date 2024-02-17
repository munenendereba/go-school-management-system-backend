package models

import "time"

type GuardianTable interface {
	TableName() string
}

func (Guardian) TableName() string {
	return "guardian"
}

type Guardian struct {
	Id              int       `gorm:"primarykey;autoIncrement"`
	FirstName       *string   `json:"firstName" gorm:"column:firstName"`
	LastName        *string   `json:"lastName" gorm:"column:lastName"`
	MiddleName      *string   `json:"middleName" gorm:"column:middleName"`
	Gender          *string   `json:"gender" gorm:"column:gender;not null;"`
	PhoneNumber     string    `json:"phoneNumber" gorm:"column:phoneNumber"`
	Email           string    `json:"email"`
	PhysicalAddress string    `json:"physicalAddress" gorm:"column:physicalAddress"`
	PostalAddress   string    `json:"postalAddress" gorm:"column:postalAddress"`
	City            string    `json:"city"`
	County          string    `json:"county"`
	Country         string    `json:"country"`
	CreatedAt       time.Time `json:"createdAt"  gorm:"column:createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"  gorm:"column:updatedAt"`
}
