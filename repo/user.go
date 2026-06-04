package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	FindUserByEmail(email string, password string) (*User, error)
	// Login(email string, password string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {

	query := `INSERT INTO users (
	first_name, 
	last_name, 
	email, 
	password, 
	is_shop_owner
	) 
	VALUES (
	:first_name, :last_name, :email, :password, :is_shop_owner
	) 
	RETURNING id`

	var userId int

	row, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, fmt.Errorf("Failed to execute query: %v", err)
	}
	if row.Next() {
		err = row.Scan(&userId)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan result: %v", err)
		}
	}

	user.ID = userId
	return &user, nil
}

func (r *userRepo) FindUserByEmail(email string, password string) (*User, error) {
	if email == "" {
		return nil, fmt.Errorf("Empty email provided")
	}
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users WHERE email = $1`
	var user User
	err := r.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found")
		}
		return nil, fmt.Errorf("Failed to execute query: %v", err)
	}
	fmt.Printf("User found: %+v\n", user)
	if user.Password == password {
		return &user, nil
	}
	return nil, fmt.Errorf("Invalid email or password")
}
