package auth_service

import (
	"Gin-Todo/models"
)

type User struct {
	NickName string
	Email    string
	Pwd      string
}

func (u *User) Find() (*models.User, error) {
	return models.GetUser(map[string]interface{}{
		"email":    u.Email,
		"password": u.Pwd,
	})
}

func (u *User) Create() error {
	uid, _ := models.GetUUID(u.Email)
	err := models.AddUser(map[string]interface{}{
		"uid":       uid,
		"nick_name": u.NickName,
		"email":     u.Email,
		"pwd":       u.Pwd,
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *User) ExistByEmail() bool {
	if models.ExistUserByEmail(u.Email) {
		return true
	}
	return false
}
