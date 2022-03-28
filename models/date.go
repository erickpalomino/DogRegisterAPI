package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Date struct {
	gorm.Model
	DateTime time.Time `json:"date"`
	Doctor   string    `json:"doctor"`
	Owner    string    `json:"owner"`
	Dog      uint      `json:"dog"`
}

func (d *Date) SaveDate() (*Date, error) {
	var err error
	err = DB.Create(&d).Error
	if err != nil {
		return &Date{}, err
	}
	return d, nil
}
