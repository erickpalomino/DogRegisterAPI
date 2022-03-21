package models

import (
	"github.com/jinzhu/gorm"
)

type Dog struct {
	gorm.Model

	DNI   string `gorm:"not null;unique" json:"dni"`
	Name  string `gorm:"size:255;not null" json:"name"`
	Race  string `gorm:"size:255;not null" json:"race"`
	Genre string `gorm:"not null" json:"genre"`
	Birth string `gorm:"not null" json:"birth"`
	Pic   string `gorm:"size:255" json:"pic"`
}

func (d *Dog) SaveDog() (*Dog, error) {
	var err error
	err = DB.Create(&d).Error
	if err != nil {
		return &Dog{}, err
	}
	return d, nil
}

func GetDogByName(name string) ([]Dog, error) {
	var d []Dog
	if err := DB.Where(&Dog{Name: name}).Find(&d); err != nil {
		return d, err.Error
	}
	return d, nil
}

func GetDogByDni(dni string) (Dog, error) {
	var d Dog
	if err := DB.Where(&Dog{DNI: dni}).Find(&d); err != nil {
		return d, err.Error
	}
	return d, nil
}
