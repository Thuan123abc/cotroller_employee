package database

import (
	"fmt"
	"gorm.io/gorm"
)

type CertificateRepo struct {
	db *gorm.DB
}

func NewCertificateRepo(db *gorm.DB) *CertificateRepo {
	db.AutoMigrate(&Certificate{})
	return &CertificateRepo{
		db: db,
	}
}

func (c CertificateRepo) CreateCerti(Certi Certificate) error {
	err := c.db.Create(&Certi).Error
	if err != nil {
		return err
	}
	return nil
}

func (c CertificateRepo) DeleteCerti(id int64) error {
	err := c.db.Delete("id_certificate", id).Error
	if err != nil {
		fmt.Println("delete fail%v\n", err)
		return err
	}
	fmt.Println("successful delete\n")
	return nil
}

func (c CertificateRepo) UpdateCerti(model Certificate) error {
	err := c.db.Model(&model).Omit("id_certificate").Updates(model).Error
	if err != nil {
		fmt.Println("update fail%v\n", err)
		return err
	}
	fmt.Println("successful update\n")
	return nil
}

func (c CertificateRepo) GetListCerti() ([]Certificate, error) {
	var models []Certificate
	err := c.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (c CertificateRepo) GetCerti(id int64) (*Certificate, error) {
	var model Certificate
	err := c.db.Where("id_certificate=?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
