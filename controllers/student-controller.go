package controllers

import (
	"errors"
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

	if res.Error != nil {
		return nil, res.Error
	}

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
	var updatedStudent models.Student

	result := db.DB.Model(&updatedStudent).Where("admissionNumber = ?", student.AdmissionNumber).Updates(student)

	if result.RowsAffected < 1 {
		if result.Error != nil {
			return nil, errors.New("student not updated")
		} else {
			return nil, nil
		}
	}

	return student, nil
}

// return bool - not found and error
func DeleteStudent(admissionNumber string) (int64, error) {
	var deletedStudent models.Student
	res := db.DB.Where("admissionNumber = ?", admissionNumber).Delete(&deletedStudent)

	if res.RowsAffected < 1 {
		//when no row deleted and no error, we assume record not found
		if res.Error == nil {
			return res.RowsAffected, nil
		} else {
			return res.RowsAffected, fmt.Errorf(res.Error.Error())
		}
	}

	return res.RowsAffected, nil
}
