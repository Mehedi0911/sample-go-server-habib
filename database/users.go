package database

import "fmt"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var userList []User

func (u User) Create() *User {
	if u.ID != 0 {
		return &u
	}
	u.ID = len(userList) + 1
	userList = append(userList, u)
	return &u
}

func (u *User) Login() {

}

func FindUserByEmail(email string, password string) (*User, error) {
	if email == "" {
		return nil, fmt.Errorf("Empty email provided")
	}
	for _, user := range userList {
		if user.Email == email {
			if user.Password == password {
				return &user, nil
			}
		}
	}
	return nil, fmt.Errorf("Invalid email or password")
}
