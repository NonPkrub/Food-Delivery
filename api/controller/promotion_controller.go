package controller

import (
	"Food-delivery/domain"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionController struct {
	promotionUseCase domain.PromotionUseCase
}

func NewPromotionController(promotionUseCase domain.PromotionUseCase) *PromotionController {
	return &PromotionController{promotionUseCase: promotionUseCase}
}

func (p *PromotionController) CreatePromotion(c *fiber.Ctx) error {
	var req *domain.PromotionForm
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.PromotionForm
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

	err = p.promotionUseCase.CreatePromotion(&newReq)
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

func (p *PromotionController) EditPromotion(c *fiber.Ctx) error {
	var req *domain.PromotionForm
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var newReq domain.PromotionForm
	err = json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return err
	}

	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
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

	err = p.promotionUseCase.EditPromotion(&newReq, uint(idInt))
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

func (p *PromotionController) DeletePromotion(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = p.promotionUseCase.DeletePromotion(uint(idInt))
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

func (p *PromotionController) GetPromotionById(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	res, err := p.promotionUseCase.GetPromotionById(uint(idInt))
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

func (p *PromotionController) GetAllPromotion(c *fiber.Ctx) error {
	res, err := p.promotionUseCase.GetAllPromotion()
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

// func (p *PromotionController) SearchPromotion(c *fiber.Ctx) error {
// 	var req *domain.Promotion
// 	jsonData, err := json.Marshal(req)
// 	if err != nil {
// 		return err
// 	}

// 	var newReq domain.Promotion
// 	err = json.Unmarshal(jsonData, &newReq)
// 	if err != nil {
// 		return err
// 	}

// 	if err := c.BodyParser(&newReq); err != nil {
// 		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
// 			"status":      fiber.ErrInternalServerError.Message,
// 			"status_code": fiber.ErrInternalServerError.Code,
// 			"message":     err.Error(),
// 			"result":      nil,
// 		})
// 	}

// 	res, err := p.promotionUseCase.SearchPromotion(&newReq)
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
// 	})

// }
