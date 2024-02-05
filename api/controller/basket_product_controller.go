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

func (bpc *BasketProductController) AddProductInBasket(c *fiber.Ctx) error {
	var form *domain.BasketProduct
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.BasketProduct
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	id := c.Params("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	newForm.BasketID = uint(idInt)

	if err := c.BodyParser(&newForm); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = bpc.basketProductUseCase.AddProductInBasket(&newForm)
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

func (bpc *BasketProductController) EditProductInBasket(c *fiber.Ctx) error {
	var form *domain.BasketProduct
	id := c.Params("id")

	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.BasketProduct
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	newForm.BasketID = uint(idInt)

	if err := c.BodyParser(&newForm); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = bpc.basketProductUseCase.EditProductInBasket(&newForm)
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

func (bpc *BasketProductController) DeleteProductInBasket(c *fiber.Ctx) error {
	var form *domain.BasketProduct
	id := c.Params("id")

	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.BasketProduct
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	newForm.BasketID = uint(idInt)

	if err := c.BodyParser(&newForm); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = bpc.basketProductUseCase.DeleteProductInBasket(&newForm)
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

// func (bpc *BasketProductController) GetProductInBasket(c *fiber.Ctx) error {
// 	var req *domain.BasketProduct
// 	id := c.Params("id")

// 	jsonData, err := json.Marshal(req)
// 	if err != nil {
// 		return err
// 	}

// 	var newReq domain.BasketProduct
// 	err = json.Unmarshal(jsonData, &newReq)
// 	if err != nil {
// 		return err
// 	}

// 	idInt, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	newReq.BasketID = uint(idInt)

// 	res, totalPrice, err := bpc.basketProductUseCase.GetProductInBasket(&newReq)
// 	if err != nil {
// 		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
// 			"status":      fiber.ErrInternalServerError.Message,
// 			"status_code": fiber.ErrInternalServerError.Code,
// 			"message":     err.Error(),
// 			"result":      nil,
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"status":      "OK",
// 		"status_code": fiber.StatusOK,
// 		"message":     "",
// 		"result":      res,
// 		"totalPrice":  totalPrice,
// 	})
// }
