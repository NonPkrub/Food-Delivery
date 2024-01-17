package controller

import (
	"Food-delivery/domain"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BasketProductController struct {
	basketProductUseCase domain.BasketProductUseCase
}

func NewBasketProductController(basketProductUseCase domain.BasketProductUseCase) *BasketProductController {
	return &BasketProductController{basketProductUseCase: basketProductUseCase}
}

func (b *BasketProductController) AddProductInBasket(c *fiber.Ctx) error {
	var req *domain.BasketProduct
	id := c.Params("id")

	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.BasketProduct
	err = json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return err
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	newReq.BasketID = uint(idInt)

	if err := c.BodyParser(&newReq); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = b.basketProductUseCase.AddProductInBasket(&newReq)
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

func (b *BasketProductController) EditProductInBasket(c *fiber.Ctx) error {
	var req *domain.BasketProduct
	id := c.Params("id")

	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.BasketProduct
	err = json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return err
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	newReq.BasketID = uint(idInt)

	if err := c.BodyParser(&newReq); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = b.basketProductUseCase.EditProductInBasket(&newReq)
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

func (b *BasketProductController) DeleteProductInBasket(c *fiber.Ctx) error {
	var req *domain.BasketProduct
	id := c.Params("id")

	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.BasketProduct
	err = json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return err
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	newReq.BasketID = uint(idInt)

	err = b.basketProductUseCase.DeleteProductInBasket(&newReq)
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

func (b *BasketProductController) GetProductInBasket(c *fiber.Ctx) error {
	var req *domain.BasketProduct
	id := c.Params("id")

	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.BasketProduct
	err = json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return err
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	newReq.BasketID = uint(idInt)

	res, totalPrice, err := b.basketProductUseCase.GetProductInBasket(&newReq)
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
		"totalPrice":  totalPrice,
	})
}
