package http

import (
	"errors"
	"strconv"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/service"
	"github.com/KaningNoppasin/Web-Server-Fiber/pkg/response"
	"github.com/KaningNoppasin/Web-Server-Fiber/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service port.ProductService
}

func NewProductHandler(service port.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAllProduct(c *fiber.Ctx) error {
	// api/products?sort=product_name-desc
	// api/products?sort=product_name-asc
	params := c.Queries()
	sort := params["sort"]
	// for sorting
	if sort != "" {
		products, err := h.service.GetAllProductSorted(sort)
		if err != nil {
			if errors.Is(err, service.ErrProductNotFound) {
				return response.SendErrorResponse(c, fiber.StatusNotFound, err)
			} else if errors.Is(err, service.ErrBadRequest) {
				return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
			}
			return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return response.SendSuccessResponse(c, products)
	}

	// api/products
	products, err := h.service.GetAllProduct()
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, products)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	productId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	product, err := h.service.GetProductByID(uint(productId))
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, product)
}

func (h *ProductHandler) GetProductByBarcode(c *fiber.Ctx) error {
	product_barcode := c.Params("barcode")
	product, err := h.service.GetProductByBarcode(product_barcode)
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, product)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product entity.Product
	// get product price from form data and convert in to integer
	productPrice, err := strconv.ParseUint(c.FormValue("price"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get product quantity from form data and convert in to integer
	quantity, err := strconv.ParseUint(c.FormValue("quantity"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// Save image to ./upload to folder
	filePath, _ := util.SaveImage(c)

	product = entity.Product{
		ProductBarcode: c.FormValue("product_barcode"),
		ProductName:    c.FormValue("product_name"),
		ImagePath:      filePath,
		Price:          uint(productPrice),
	}

	err = h.service.CreateProduct(&product, uint(quantity))
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var product entity.Product
	productId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	productPrice, err := strconv.ParseUint(c.FormValue("price"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get current product for delete image
	filePath, _ := util.SaveImage(c)

	product = entity.Product{
		ProductBarcode: c.FormValue("product_barcode"),
		ProductName:    c.FormValue("product_name"),
		ImagePath:      filePath,
		Price:          uint(productPrice),
	}
	product.ID = uint(productId)

	err = h.service.UpdateProduct(&product)
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = h.service.DeleteProduct(uint(productId))
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
