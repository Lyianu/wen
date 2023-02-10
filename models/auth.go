package models

import (
	"errors"

	"github.com/Lyianu/wen/util"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	hashed, err := GetHashedPassword(username)
	//fmt.Printf("Validating USER: %q, Pass: %q, Expected: %q", username, password, hashed)
	if err == nil {
		return util.ValidatePassword(hashed, password)
	}

	return false
}

func AddAuth(username, password string) bool {
	hashed, err := util.HashPassword(password)
	//fmt.Printf("Created USER: %q, Passwd: %q, PasswdHash: %q", username, password, hashed)
	if err != nil {
		return false
	}
	if !CheckAuth(username, password) {
		db.Create(&User{
			Username: username,
			Password: hashed,
		})
	}
	return true
}

func GetHashedPassword(username string) (string, error) {
	var auth User
	db.Select("id", "password").Where(User{Username: username}).First(&auth)
	if auth.ID > 0 {
		return auth.Password, nil
	}
	return "", errors.New("USER_NOT_EXISTS")
}
