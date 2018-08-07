package models

type User struct {
	Model

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
