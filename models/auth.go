package models

import "errors"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth User
	db.Select("id").Where(User{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

func AddAuth(username, password string) bool {
	if !CheckAuth(username, password) {
		db.Create(&User{
			Username: username,
			Password: password,
		})
	}
	return true
}

func GetHashedPassword(username string) (string, error) {
	var auth User
	db.Select("id").Where(User{Username: username}).First(&auth)
	if auth.ID > 0 {
		return auth.Password, nil
	}
	return "", errors.New("USER_NOT_EXISTS")
}
