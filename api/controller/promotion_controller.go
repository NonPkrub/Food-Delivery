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

func (pc *PromotionController) CreatePromotion(c *fiber.Ctx) error {
	var form *domain.PromotionForm
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.PromotionForm
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

	err = pc.promotionUseCase.CreatePromotion(&newForm)
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

func (pc *PromotionController) EditPromotion(c *fiber.Ctx) error {
	var form *domain.PromotionForm
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.PromotionForm
	err = json.Unmarshal(jsonData, &newForm)
	if err != nil {
		return err
	}

	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
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

	err = pc.promotionUseCase.EditPromotion(&newForm, uint(idInt))
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

func (pc *PromotionController) DeletePromotion(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = pc.promotionUseCase.DeletePromotion(uint(idInt))
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

func (pc *PromotionController) GetPromotionById(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	result, err := pc.promotionUseCase.GetPromotionById(uint(idInt))
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

func (pc *PromotionController) GetAllPromotion(c *fiber.Ctx) error {
	queryCode := c.Query("code")
	queryName := c.Query("name")

	result, err := pc.promotionUseCase.GetAllPromotion(queryCode, queryName)
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
