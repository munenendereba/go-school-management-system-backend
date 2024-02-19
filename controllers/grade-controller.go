package controllers

import (
	"munenendereba/sms-backend/db"
	"munenendereba/sms-backend/models"

	"github.com/jinzhu/gorm"
)

func GetGrades() ([]*models.Grade, error) {
	var grades []*models.Grade

	res := db.DB.Find(&grades)

	if res.Error != nil {
		return nil, res.Error
	}

	return grades, nil
}

func GetGrade(id string) (*models.Grade, error) {
	var grade models.Grade

	res := db.DB.First(&grade, "id = ?", id)
	if res.Error != nil {
		if gorm.IsRecordNotFoundError(res.Error) {
			return nil, nil
		}

		return nil, res.Error
	}

	return &grade, nil
}

func CreateGrade(grade *models.Grade) (*models.Grade, error) {
	res := db.DB.Create(&grade)

	if res.Error != nil {
		return nil, res.Error
	}

	return grade, nil
}

func UpdateGrade(grade *models.Grade) (*models.Grade, error) {
	var updatedGrade models.Grade

	result := db.DB.Model(&updatedGrade).Where("id = ?", grade.Id).Updates(grade)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected < 1 {
		return nil, nil
	}

	return &updatedGrade, nil
}

func DeleteGrade(id string) (int64, error) {
	var deletedGrade models.Grade

	res := db.DB.Where("id = ?", id).Delete(&deletedGrade)

	if res.Error != nil {
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}
