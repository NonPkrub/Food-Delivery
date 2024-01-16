package controller

import (
	"Food-delivery/domain"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// import (
// 	"Food-delivery/domain"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// )

type BasketController struct {
	basketUseCase domain.BasketUseCase
}

func NewBasketController(basketUseCase domain.BasketUseCase) *BasketController {
	return &BasketController{basketUseCase: basketUseCase}
}

func (b *BasketController) CreateBasket(c *fiber.Ctx) error {
	var req domain.BasketForm
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	req.UserID = uint(idInt)

	err = b.basketUseCase.CreateBasket(&req)
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
		"result":      nil,
	})
}

func (b *BasketController) AddPromotionBasket(c *fiber.Ctx) error {
	var req *domain.BasketPromotionForm
	id := c.Params("id")

	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.BasketPromotionForm
	err = json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return err
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	newReq.UserID = uint(idInt)

	if err := c.BodyParser(&newReq); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = b.basketUseCase.AddPromotionBasket(&newReq)
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
		"result":      nil,
	})
}

func (b *BasketController) DeletePromotionBasket(c *fiber.Ctx) error {

	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = b.basketUseCase.DeletePromotionBasket(uint(idInt))
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
		"result":      nil,
	})
}

func (b *BasketController) GetBasketByUserId(c *fiber.Ctx) error {

	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	basket, err := b.basketUseCase.GetBasketByUserId(uint(idInt))
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
		"result":      basket,
	})
}
