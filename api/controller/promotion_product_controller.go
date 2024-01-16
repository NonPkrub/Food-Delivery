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

func (p *PromotionProductController) AddPromotionProduct(c *fiber.Ctx) error {
	var req *domain.PromotionProductForm
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.PromotionProductForm
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

	err = p.promotionProductUseCase.AddPromotionProduct(&newReq)
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

func (p *PromotionProductController) EditPromotionProduct(c *fiber.Ctx) error {
	var req domain.PromotionProductForm
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	req.PromotionID = uint(idInt)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	err = p.promotionProductUseCase.EditPromotionProduct(&req)
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

func (p *PromotionProductController) GetPromotionProduct(c *fiber.Ctx) error {
	var req domain.PromotionProductForm
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	req.PromotionID = uint(idInt)

	// if err := c.BodyParser(&req); err != nil {
	// 	return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
	// 		"status":      fiber.ErrInternalServerError.Message,
	// 		"status_code": fiber.ErrInternalServerError.Code,
	// 		"message":     err.Error(),
	// 		"result":      nil,
	// 	})
	// }

	res, err := p.promotionProductUseCase.GetPromotionProduct(&req)
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
