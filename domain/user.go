package domain

import "github.com/golang-jwt/jwt/v4"

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
	SignUp(u *UserSignUpForm) (*UserReply, error)
	Login(u *UserLoginForm) (*TokenReply, error)
	GetUserById(id uint) (*UserReply, error)
}

type UserRepository interface {
	SignUp(u *User) (*UserReply, error)
	Login(u *User) (*UserLoginReply, error)
	GetUserById(u *User, id uint) (*UserReply, error)
}
