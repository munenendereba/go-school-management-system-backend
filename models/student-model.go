package models

import (
	"time"
)

type Tabler interface {
	TableName() string
}

func (Student) TableName() string {
	return "student"
}

type Student struct {
	AdmissionNumber  string    `json:"admissionNumber" gorm:"primarykey;column:admissionNumber"`
	FirstName        string    `json:"firstName" gorm:"column:firstName"`
	LastName         string    `json:"lastName" gorm:"column:lastName"`
	MiddleName       string    `json:"middleName" gorm:"column:middleName"`
	BirthCertificate string    `json:"birthCertificate" gorm:"column:birthCertificate"`
	NemisNumber      string    `json:"nemisNumber" gorm:"column:nemisNumber"`
	Gender           string    `json:"gender"`
	DateOfBirth      time.Time `json:"dateOfBirth" gorm:"column:dateOfBirth"`
	Status           string    `json:"status"`
	DateJoined       time.Time `json:"dateJoined" gorm:"column:dateJoined"`
	Type             string    `json:"type"`
	PhoneNumber      string    `json:"phoneNumber" gorm:"column:phoneNumber"`
	Email            string    `json:"email"`
	PhysicalAddress  string    `json:"physicalAddress" gorm:"column:physicalAddress"`
	PostalAddress    string    `json:"postalAddress" gorm:"column:postalAddress"`
	City             string    `json:"city"`
	County           string    `json:"county"`
	Country          string    `json:"country"`
	CreatedAt        time.Time `json:"createdAt"  gorm:"column:createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"  gorm:"column:updatedAt"`
}
