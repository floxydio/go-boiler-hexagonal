package http

import (
	"strconv"

	"github.com/floxydio/boiler-fiber/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service domain.ProductService
}

func NewProductHandler(service domain.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/product", h.GetAll)
	app.Post("/product/create", h.Create)
	app.Get("/product/detail/:id_product", h.GetProductById)
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.service.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"error":   true,
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"error":   false,
		"data":    data,
		"message": "Successfully Get Product",
	})

}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	data := domain.ProductForm{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	errCreate := h.service.Create(&data)

	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"error":   true,
			"message": "Internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  201,
		"error":   false,
		"message": "Success!",
	})

}

func (h *ProductHandler) GetProductById(c *fiber.Ctx) error {
	param := c.Params("id_product")

	idProduct, _ := strconv.ParseUint(param, 10, 64)

	data, err := h.service.FindById(idProduct)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"error":   true,
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"error":   false,
		"data":    data,
		"message": "Successfully Get Product",
	})
}
