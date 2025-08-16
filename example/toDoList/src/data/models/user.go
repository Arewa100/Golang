package models

type User struct {
	UserName string
	Password string
}

func (user *User) GetUsrName() string {
	return user.UserName
}

func (user *User) GetPassword() string {
	return user.Password
}
