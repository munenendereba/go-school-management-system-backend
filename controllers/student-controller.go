package controllers

import (
	"munenendereba/sms-backend/db"
	models "munenendereba/sms-backend/models"
)

func GetStudents() ([]*models.Student, error) {
	var students []*models.Student

	res := db.DB.Find(&students)

	if res.Error != nil {
		return nil, res.Error
	}

	return students, nil
}
