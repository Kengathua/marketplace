package models

import (
	"time"

	"github.com/matawis/matawis/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	common.BioData `gorm:"embedded"`
	GUID           *string   `json:"guid"`
	Password       string    `json:"password"`
	LastLogin      time.Time `json:"last_login"`
	IsStaff        bool      `gorm:"column:is_staff" json:"is_staff"`
	IsAdmin        bool      `gorm:"column:is_admin" json:"is_admin"`
	IsSuperUser    bool      `gorm:"column:is_superuser" json:"is_superuser"`
}

func (u *User) GeneratePasswordHarsh() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
	return err
}

func (u *User) CheckPasswordHarsh(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

var user = User{Password: "Password"}
var err = user.GeneratePasswordHarsh()
