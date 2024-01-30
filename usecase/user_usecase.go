package usecase

import (
	"Food-delivery/domain"
	"fmt"
	"net/mail"
	"os"
	"regexp"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

func (uc *userUseCase) SignUp(form *domain.UserSignUpForm) (*domain.UserReply, error) {

	var err error
	if len(form.Password) < 10 {
		return nil, err
	}

	reCheckPassword := regexp.MustCompile("[^a-zA-Z0-9!@#$%^&*()_+]+")
	password := reCheckPassword.ReplaceAllString(form.Password, "")

	if !containsUppercase(password) {
		return nil, err
	}

	if !containsLowercase(password) {
		return nil, err
	}

	_, err = mail.ParseAddress(form.Email)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	form.Password = string(hashed)

	users := &domain.User{
		FirstName:     form.FirstName,
		LastName:      form.LastName,
		Email:         form.Email,
		Password:      form.Password,
		PhoneNumber:   form.PhoneNumber,
		NationalID:    form.NationalID,
		Address:       form.Address,
		DetailAddress: form.DetailAddress,
	}

	user, err := uc.userRepo.CreateUser(users)
	if err != nil {
		return nil, err
	}

	userReply := &domain.UserReply{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		NationalID:    user.NationalID,
		Address:       user.Address,
		DetailAddress: user.DetailAddress,
	}

	return userReply, nil

}

func (uc *userUseCase) Login(form *domain.UserLoginForm) (*domain.TokenReply, error) {

	var err error
	if len(form.Password) < 10 {
		return nil, err
	}
	re := regexp.MustCompile("[^a-zA-Z0-9!@#$%^&*()_+]+")
	password := re.ReplaceAllString(form.Password, "")

	if !containsUppercase(password) {
		return nil, err
	}

	if !containsLowercase(password) {
		return nil, err
	}

	_, err = mail.ParseAddress(form.Email)
	if err != nil {
		return nil, err
	}

	users := &domain.User{
		Email:    form.Email,
		Password: form.Password,
	}

	user, err := uc.userRepo.FindOne(users)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		return nil, err
	}

	// claims := jwt.StandardClaims{
	// 	Issuer:    users.Email,
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	// }

	claims := domain.UsersClaims{
		Email: form.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	res := &domain.TokenReply{
		AccessToken: token,
	}

	return res, nil

}

func (uc *userUseCase) GetUserByID(id uint) (*domain.UserReply, error) {
	var user domain.User
	users, err := uc.userRepo.GetOneByID(&user)
	if err != nil {
		return nil, err
	}

	userReply := &domain.UserReply{
		FirstName:     users.FirstName,
		LastName:      users.LastName,
		Email:         users.Email,
		PhoneNumber:   users.PhoneNumber,
		NationalID:    users.NationalID,
		Address:       users.Address,
		DetailAddress: users.DetailAddress,
	}

	return userReply, nil
}

func (uc *userUseCase) GetMe() (*domain.UserReply, error) {
	var user domain.User
	users, err := uc.userRepo.GetOneByID(&user)
	if err != nil {
		return nil, err
	}

	userReply := &domain.UserReply{
		FirstName:     users.FirstName,
		LastName:      users.LastName,
		Email:         users.Email,
		PhoneNumber:   users.PhoneNumber,
		NationalID:    users.NationalID,
		Address:       users.Address,
		DetailAddress: users.DetailAddress,
	}

	return userReply, nil
}
