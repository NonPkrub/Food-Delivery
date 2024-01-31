package controller

import (
	"Food-delivery/domain"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productUseCase domain.ProductUseCase
}

func NewProductController(productUseCase domain.ProductUseCase) *ProductController {
	return &ProductController{productUseCase: productUseCase}
}

func (pc *ProductController) GetAll(c *fiber.Ctx) error {
	result, err := pc.productUseCase.GetAll()
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

func (pc *ProductController) AddProduct(c *fiber.Ctx) error {
	var form *domain.ProductForm
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.ProductForm
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

	result, err := pc.productUseCase.AddProduct(&newForm)
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

func (pc *ProductController) EditProduct(c *fiber.Ctx) error {
	var form *domain.ProductForm
	jsonData, err := json.Marshal(form)
	if err != nil {
		return err
	}

	var newForm domain.ProductForm
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

	result, err := pc.productUseCase.EditProduct(&newForm, uint(idInt))
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

func (pc *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = pc.productUseCase.DeleteProduct(uint(idInt))
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
