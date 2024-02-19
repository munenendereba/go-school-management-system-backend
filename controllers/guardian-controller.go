package controllers

import (
	"munenendereba/sms-backend/db"
	"munenendereba/sms-backend/models"

	"github.com/jinzhu/gorm"
)

func GetGuardians() ([]*models.Guardian, error) {
	var guardians []*models.Guardian

	res := db.DB.Find(&guardians)

	if res.Error != nil {
		return nil, res.Error
	}

	return guardians, nil
}

func GetGuardian(id string) (*models.Guardian, error) {
	var guardian models.Guardian

	res := db.DB.First(&guardian, "id = ?", id)

	if res.Error != nil {
		if gorm.IsRecordNotFoundError(res.Error) {
			return nil, nil
		}

		return nil, res.Error
	}

	if res.RowsAffected < 1 {
		return nil, nil
	}

	return &guardian, nil
}

func CreateGuardian(guardian *models.Guardian) (*models.Guardian, error) {
	res := db.DB.Create(&guardian)

	if res.Error != nil {
		return nil, res.Error
	}

	return guardian, nil
}

func UpdateGuardian(guardian *models.Guardian) (*models.Guardian, error) {
	var updatedGuardian models.Guardian

	result := db.DB.Model(&updatedGuardian).Where("id = ?", guardian.Id).Updates(guardian)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected < 1 {
		return nil, nil
	}

	return &updatedGuardian, nil
}

func DeleteGuardian(id string) (int64, error) {
	var deletedGuardian models.Guardian

	res := db.DB.Where("id = ?", id).Delete(&deletedGuardian)

	if res.Error != nil {
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}
