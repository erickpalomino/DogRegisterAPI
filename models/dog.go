package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Dog struct {
	gorm.Model
	DNI         uint64       `gorm:"not null;unique" json:"dni"`
	Name        string       `gorm:"size:255;not null" json:"name"`
	Race        string       `gorm:"size:255;not null" json:"race"`
	Genre       string       `gorm:"not null" json:"genre"`
	Birth       time.Time    `gorm:"not null" json:"birth"`
	Pic         string       `gorm:"size:255" json:"pic"`
	Diagnostics []Diagnostic `json:"diagnostics"`
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

func GetDogByDni(dni uint64) (Dog, error) {
	var d Dog
	if err := DB.Where(&Dog{DNI: dni}).Find(&d); err != nil {
		return d, err.Error
	}
	return d, nil
}

func (d *Dog) GetDiagnosticsFromDog() []Diagnostic {
	var diagnostics []Diagnostic
	if err := DB.Where("dog_id =?", d.ID).Find(&diagnostics).Error; err != nil {
		fmt.Print(err.Error())
	}
	return diagnostics
}
