package models

import (
	"errors"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	hashids "github.com/speps/go-hashids"
)

type User struct {
	Model

	Uid      string `json:"uid"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Password string
}

func CheckUser(email, password string) bool {
	var user User
	db.Select("id").Where(User{Email: email, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func AddUser(maps map[string]interface{}) error {
	user := User{
		Uid:      maps["uid"].(string),
		NickName: maps["nick_name"].(string),
		Email:    maps["email"].(string),
		Password: maps["pwd"].(string),
	}
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(maps map[string]interface{}) (*User, error) {
	user := User{}
	err := db.Where(maps).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func EditUser(id int, data interface{}) bool {
	db.Model(&User{}).Where("id = ?", id).Update(data)
	return true
}

func GetAuthenticateUser(authenticate AuthenticateRequest) (User, error) {
	var users []User
	db.Where(authenticate).Find(&users)
	switch len(users) {
	case 0:
		return User{}, errors.New("User not found")
	case 1:
		return users[0], nil
	default:
		return User{}, errors.New("Internal error: Several users found")
	}
}

func ExistUserByEmail(email string) bool {
	users := []User{}
	db.Debug().Where("email = ? ", email).Find(&users)
	if len(users) > 0 {
		return true
	}
	return false
}

func ExistUserByID(id int) bool {
	var tag User
	db.Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func GetUUID(email string) (string, error) {
	hdata := hashids.NewData()
	hdata.MinLength = 6
	hdata.Salt = time.Now().String() + email // 设置秘钥
	hid, _ := hashids.NewWithData(hdata)
	numbers := []int{20, 199, 260, 99}
	hash, err := hid.Encode(numbers)
	if err != nil {
		return "", err
	}

	return hash, nil
}
