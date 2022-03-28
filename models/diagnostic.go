package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Diagnostic struct {
	gorm.Model

	Symptom     string    `json:"symptom"`
	Medicines   string    `json:"medicines"`
	Price       float32   `json:"price"`
	BloodResult string    `json:"bloodResult"`
	XrayPic     string    `json:"xrayPic"`
	Diagnostic  string    `json:"diagnostic"`
	Date        time.Time `json:"date"`
	Doctor      string    `json:"doctor""`
	DogID       uint      `json:"dogID"`
}

func (d *Diagnostic) SaveDiagnostic() (*Diagnostic, error) {
	var err error
	err = DB.Create(&d).Error
	if err != nil {
		return &Diagnostic{}, err
	}
	return d, nil
}

func GetDiagnosticById(id uint64) (Diagnostic, error) {
	var d Diagnostic
	if err := DB.Find(&d, id).Error; err != nil {
		return d, err
	}
	return d, nil
}
