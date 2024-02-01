package usecase

import (
	"Food-delivery/domain"
	"errors"
	"fmt"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo      domain.UserRepository
	basketUseCase domain.BasketUseCase
}

func NewUserUseCase(userRepo domain.UserRepository, basketUseCase domain.BasketUseCase) domain.UserUseCase {
	return &userUseCase{userRepo: userRepo, basketUseCase: basketUseCase}
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

func IsValidPassword(form *domain.UserLoginForm) error {
	if len(form.Password) < 10 {
		return errors.New("Password must be at least 10 characters")
	}

	nonASCII := regexp.MustCompile("[^\x00-\x7F]+")
	if nonASCII.MatchString(form.Password) {
		return errors.New("Password accept only English character")
	}

	password := form.Password

	if !containsUppercase(password) {
		return errors.New("Password must contain the Upper case at least one character")
	}

	if !containsLowercase(password) {
		return errors.New("Password must contain Lower case at least one character")
	}

	_, err := mail.ParseAddress(form.Email)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) SignUp(form *domain.UserSignUpForm) (*domain.UserReply, error) {
	forms := &domain.UserLoginForm{
		Email:    form.Email,
		Password: form.Password,
	}

	email := &domain.User{
		Email:    form.Email,
		Password: form.Password,
	}

	validEmail, err := uc.userRepo.FindOne(email)
	if validEmail.Email == email.Email {
		return nil, err
	}

	err = IsValidPassword(forms)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)
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

	basket := &domain.BasketForm{
		UserID: user.ID,
	}

	err = uc.basketUseCase.CreateBasket(basket)
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
	users := &domain.User{
		Email:    form.Email,
		Password: form.Password,
	}

	forms := &domain.UserLoginForm{
		Email:    form.Email,
		Password: form.Password,
	}

	err := IsValidPassword(forms)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindOne(users)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		return nil, err
	}

	claims := domain.UsersClaims{
		Email: form.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   strconv.Itoa(int(user.ID)),
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

func (uc *userUseCase) Me(myToken string) (*domain.UserReply, error) {
	fmt.Println(myToken)
	token, err := jwt.ParseWithClaims(myToken, &domain.UsersClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims := token.Claims.(*domain.UsersClaims)

	userID := claims.RegisteredClaims.Subject
	user, err := uc.userRepo.GetMe(userID)
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
