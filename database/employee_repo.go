package database

import (
	"fmt"
	"gorm.io/gorm"
)

type EmployeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) *EmployeeRepo {
	db.AutoMigrate(&Employee{})
	return &EmployeeRepo{
		db: db,
	}
}

func (e EmployeeRepo) CreateEmployee(Employee Employee) error {
	err := e.db.Create(&Employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeeRepo) DeleteEmployee(id int64) error {
	err := e.db.Delete("id", id).Error
	if err != nil {
		fmt.Println("delete fail%v\n", err)
		return err
	}
	fmt.Println("successful delete\n")
	return nil
}

func (e EmployeeRepo) UpdateEmployee(model Employee) error {
	err := e.db.Model(&model).Omit("id").Updates(model).Error
	if err != nil {
		fmt.Println("update fail%v\n", err)
		return err
	}
	fmt.Println("successful update\n")
	return nil
}

func (e EmployeeRepo) GetListCerti() ([]Employee, error) {
	var models []Employee
	err := e.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (e EmployeeRepo) GetEmployee(id int64) (*Employee, error) {
	var model Employee
	err := e.db.Where("id?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
