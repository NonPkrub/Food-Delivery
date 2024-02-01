package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Model
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	PhoneNumber   string `json:"phone_number"`
	NationalID    string `json:"national_id"`
	Address       string `json:"address"`
	DetailAddress string `json:"detail_address"`
}

type UserSignUpForm struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	PhoneNumber   string `json:"phone_number"`
	NationalID    string `json:"national_id"`
	Address       string `json:"address"`
	DetailAddress string `json:"detail_address"`
}

type UserLoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserReply struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	NationalID    string `json:"national_id"`
	Address       string `json:"address"`
	DetailAddress string `json:"detail_address"`
}

type UserLoginReply struct {
	Password string `json:"password"`
}

type TokenReply struct {
	AccessToken string `json:"access_token"`
}

func (u *User) TableName() string {
	return "users"
}

type UsersClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type UserUseCase interface {
	SignUp(user *UserSignUpForm) (*UserReply, error)
	Login(user *UserLoginForm) (*TokenReply, error)
	GetUserByID(id uint) (*UserReply, error)
	Me(string) (*UserReply, error)
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	FindOne(user *User) (*User, error)
	GetOneByID(user *User) (*User, error)
	GetMe(string) (*User, error)
}
