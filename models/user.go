package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"dog-app/utils/token"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Type     string `gorm:"not null" json:"type"`
	Dates    []Date `json:"dates"`
}

func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}
func (u *User) PrepareGive() {
	u.Password = ""
}
func GetDoctors() (doctors []User, err error) {
	if err := DB.Where("type=?", "doctor").Find(&doctors).Error; err != nil {
		fmt.Print(err.Error())
	}
	return
}

func (doctor *User) GetDoctorDates() (dates []Date, err error) {
	if err := DB.Where("doctor=?", doctor.Username).Find(&dates).Error; err != nil {
		fmt.Print(err.Error())
	}
	return
}
func (owner *User) GetOwnerDates() (dates []Date, err error) {
	if err := DB.Where("owner=?", owner.Username).Find(&dates).Error; err != nil {
		fmt.Print(err.Error())
	}
	return
}
