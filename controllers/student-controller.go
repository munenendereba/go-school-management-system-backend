package controllers

import (
	"fmt"
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

func GetStudent(admissionNumber string) (*models.Student, error) {
	var student models.Student

	res := db.DB.First(&student, "admissionNumber = ?", admissionNumber)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("student with admission %s not found", admissionNumber)
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
	res := db.DB.Save(student)
	if res.Error != nil {
		return nil, res.Error
	}
	return student, nil
}

func DeleteStudent(admissionNumber string) error {
	var deletedStudent models.Student

	res := db.DB.Delete(&deletedStudent, "admissionNumber =?", admissionNumber)

	if res.RowsAffected == 0 {
		return fmt.Errorf("student with admission %s not found", admissionNumber)
	}
	return nil
}
