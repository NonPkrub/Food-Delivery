package usecase

import (
	"Food-delivery/domain"
	"encoding/json"
	"fmt"
	"net/mail"
	"os"
	"regexp"
	"time"
	"unicode"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func containsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

func containsLowercase(s string) bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		}
	}

	return false
}

func (u *userUseCase) SignUp(d *domain.UserSignUpForm) (*domain.UserReply, error) {

	var err error
	if len(d.Password) < 10 {
		return nil, err
	}

	re := regexp.MustCompile("[^a-zA-Z0-9!@#$%^&*()_+]+")
	password := re.ReplaceAllString(d.Password, "")

	if !containsUppercase(password) {
		return nil, err
	}

	if !containsLowercase(password) {
		return nil, err
	}

	_, err = mail.ParseAddress(d.Email)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	d.Password = string(hashed)

	users := &domain.User{
		FirstName:     d.FirstName,
		LastName:      d.LastName,
		Email:         d.Email,
		Password:      d.Password,
		PhoneNumber:   d.PhoneNumber,
		NationalID:    d.NationalID,
		Address:       d.Address,
		DetailAddress: d.DetailAddress,
	}

	user, err := u.userRepo.SignUp(users)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u *userUseCase) Login(d *domain.UserLoginForm) (*domain.TokenReply, error) {

	var err error
	if len(d.Password) < 10 {
		return nil, err
	}
	re := regexp.MustCompile("[^a-zA-Z0-9!@#$%^&*()_+]+")
	password := re.ReplaceAllString(d.Password, "")

	if !containsUppercase(password) {
		return nil, err
	}

	if !containsLowercase(password) {
		return nil, err
	}

	_, err = mail.ParseAddress(d.Email)
	if err != nil {
		return nil, err
	}
	users := &domain.User{
		Email:    d.Email,
		Password: d.Password,
	}

	user, err := u.userRepo.Login(users)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(d.Password))
	if err != nil {
		return nil, err
	}

	claims := jwt.StandardClaims{
		Issuer:    users.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	response := map[string]interface{}{
		"jwtToken": token,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	res := &domain.TokenReply{
		Token: jsonResponse,
	}

	return res, nil

}

func (u *userUseCase) GetUserById(id uint) (*domain.UserReply, error) {
	var d domain.User
	user, err := u.userRepo.GetUserById(&d, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
