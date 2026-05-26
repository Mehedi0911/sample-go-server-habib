package repo

import "fmt"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	FindUserByEmail(email string, password string) (*User, error)
	// Login(email string, password string) (*User, error)
}

type userRepo struct {
	userList []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}

}

func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}
	user.ID = len(r.userList) + 1
	r.userList = append(r.userList, user)
	return &user, nil
}

func (r *userRepo) FindUserByEmail(email string, password string) (*User, error) {
	if email == "" {
		return nil, fmt.Errorf("Empty email provided")
	}
	for _, user := range r.userList {
		if user.Email == email {
			if user.Password == password {
				return &user, nil
			}
		}
	}
	return nil, fmt.Errorf("Invalid email or password")
}
