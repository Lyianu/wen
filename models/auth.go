package models

type Auth struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

func AddAuth(username, password string) bool {
	if !CheckAuth(username, password) {
		db.Create(&Auth{
			Username: username,
			Password: password,
		})
	}
	return true
}
