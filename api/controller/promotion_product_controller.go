package controller

import (
	"Food-delivery/domain"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionProductController struct {
	promotionProductUseCase domain.PromotionProductUseCase
}

func NewPromotionProductController(promotionProductUseCase domain.PromotionProductUseCase) *PromotionProductController {
	return &PromotionProductController{promotionProductUseCase: promotionProductUseCase}
}

func (ppc *PromotionProductController) AddPromotionProduct(c *fiber.Ctx) error {
	var form *domain.PromotionProductForm
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.PromotionProductForm
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

	err = ppc.promotionProductUseCase.AddPromotionProduct(&newForm)
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

func (ppc *PromotionProductController) EditPromotionProduct(c *fiber.Ctx) error {
	var form domain.PromotionProductForm
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	form.PromotionID = uint(idInt)

	if err := c.BodyParser(&form); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = ppc.promotionProductUseCase.EditPromotionProduct(&form)
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

func (ppc *PromotionProductController) GetPromotionProduct(c *fiber.Ctx) error {
	var form domain.PromotionProductForm
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	form.PromotionID = uint(idInt)

	// if err := c.BodyParser(&req); err != nil {
	// 	return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
	// 		"status":      fiber.ErrInternalServerError.Message,
	// 		"status_code": fiber.ErrInternalServerError.Code,
	// 		"message":     err.Error(),
	// 		"result":      nil,
	// 	})
	// }

	result, err := ppc.promotionProductUseCase.GetPromotionProduct(&form)
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
