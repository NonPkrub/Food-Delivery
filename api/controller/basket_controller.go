package controller

import (
	"Food-delivery/api/middleware"
	"Food-delivery/domain"
	"encoding/json"
	"fmt"
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

// func (b *BasketController) CreateBasket(c *fiber.Ctx) error {
// 	var req domain.BasketForm
// 	id := c.Params("id")

// 	idInt, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	req.UserID = uint(idInt)

// 	err = b.basketUseCase.CreateBasket(&req)
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
// 		"result":      nil,
// 	})
// }

func (bc *BasketController) AddPromotionBasket(c *fiber.Ctx) error {
	var form *domain.BasketPromotionForm

	ID, err := middleware.UserClaim(c)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.BasketPromotionForm
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	uintID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return err
	}

	newForm.UserID = uint(uintID)

	if err := c.BodyParser(&newForm); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = bc.basketUseCase.AddPromotionBasket(&newForm)
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

func (bc *BasketController) DeletePromotionBasket(c *fiber.Ctx) error {
	ID, err := middleware.UserClaim(c)
	if err != nil {
		return err
	}

	uintID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return err
	}

	err = bc.basketUseCase.DeletePromotionBasket(uint(uintID))
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

func (bc *BasketController) GetBasketByUserId(c *fiber.Ctx) error {
	ID, err := middleware.UserClaim(c)
	fmt.Println(ID)
	if err != nil {
		return err
	}

	uintID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return err
	}

	basket, err := bc.basketUseCase.GetBasketByUserId(uint(uintID))
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
