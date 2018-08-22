package models

import (
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

func ExistUserByEmail(email string) bool {
	users := []User{}
	db.Debug().Where("email = ? ", email).Find(&users)
	if len(users) > 0 {
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
