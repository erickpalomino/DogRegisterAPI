package models

import "github.com/jinzhu/gorm"

type Diagnostic struct {
	gorm.Model
	DogDNI uint    `json:"dogDni"`
	Price  float32 `json:"price"`
}

func (d *Diagnostic) SaveDiagnostic() (*Diagnostic, error) {
	var err error
	err = DB.Create(&d).Error
	if err != nil {
		return &Diagnostic{}, err
	}
	return d, nil
}
