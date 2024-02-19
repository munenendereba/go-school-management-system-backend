package controllers

import (
	"munenendereba/sms-backend/db"
	models "munenendereba/sms-backend/models"

	"github.com/jinzhu/gorm"
)

func GetStudents() ([]*models.Student, error) {
	var students []*models.Student

	res := db.DB.Find(&students)

	if res.Error != nil {
		return nil, res.Error
	}

	return students, nil
}

func GetStudent(admissionNumber string) (*models.Student, error) {
	var student models.Student

	res := db.DB.First(&student, "admissionNumber = ?", admissionNumber)

	if res.Error != nil {
		if gorm.IsRecordNotFoundError(res.Error) {
			return nil, nil
		}

		return nil, res.Error
	}

	return &student, nil
}

func CreateStudent(student *models.Student) (*models.Student, error) {
	res := db.DB.Create(&student)
	if res.Error != nil {
		return nil, res.Error
	}
	return student, nil
}

func UpdateStudent(student *models.Student) (*models.Student, error) {
	var updatedStudent models.Student

	result := db.DB.Model(&updatedStudent).Where("admissionNumber = ?", student.AdmissionNumber).Updates(student)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected < 1 {
		return nil, nil
	}

	return &updatedStudent, nil
}

func DeleteStudent(admissionNumber string) (int64, error) {
	var deletedStudent models.Student
	res := db.DB.Where("admissionNumber = ?", admissionNumber).Delete(&deletedStudent)

	if res.Error != nil {
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}
