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

func (u *UserController) SignUp(c *fiber.Ctx) error {
	var req *domain.UserSignUpForm
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.UserSignUpForm
	err = json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&newReq); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := u.userUseCase.SignUp(&newReq)
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
		"result":      res,
	})
}

func (u *UserController) Login(c *fiber.Ctx) error {
	var req *domain.UserLoginForm
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.UserLoginForm
	err = json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&newReq); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := u.userUseCase.Login(&newReq)
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
		"result":      res,
	})
}

func (u *UserController) GetUserById(c *fiber.Ctx) error {
	user := c.Params("id")

	userInt, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		return err
	}
	fmt.Println(userInt)
	res, err := u.userUseCase.GetUserByID(uint(userInt))

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
		"result":      res,
	})
}
