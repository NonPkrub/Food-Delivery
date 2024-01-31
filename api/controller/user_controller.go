package controller

import (
	"Food-delivery/domain"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUseCase domain.UserUseCase
}

func NewUserController(userUseCase domain.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (uc *UserController) SignUp(c *fiber.Ctx) error {
	var form *domain.UserSignUpForm
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.UserSignUpForm
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&newForm); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	result, err := uc.userUseCase.SignUp(&newForm)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      result,
	})
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	var form *domain.UserLoginForm
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.UserLoginForm
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&newForm); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	result, err := uc.userUseCase.Login(&newForm)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      result,
	})
}

func (uc *UserController) GetUserById(c *fiber.Ctx) error {
	user := c.Params("id")

	userInt, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		return err
	}
	fmt.Println(userInt)
	result, err := uc.userUseCase.GetUserByID(uint(userInt))

	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      result,
	})
}

func (uc *UserController) Me(c *fiber.Ctx) error {
	token := c.Query("token")

	result, err := uc.userUseCase.Me(token)

	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      result,
	})
}
