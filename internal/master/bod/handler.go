package bod

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/asastarealty-backend/app/response"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type BodHandler struct {
	bodService BodService
}

func NewBodHandler(router fiber.Router, bs BodService) {
	handler := &BodHandler{bodService: bs}
	router.Get("/bod", handler.GetAll)
	router.Get("/bod/:id", handler.GetById)
}

func (bh *BodHandler) GetAll(c *fiber.Ctx) error {
	listBod, err := bh.bodService.FindAll()
	log.Println("Entering BOD Handler GetAll Error....")
	log.Printf("rows Handler -> %v", listBod)
	log.Printf("err Handler  -> %v", err)
	if err != nil {
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "healthty", nil)
}

func (bh *BodHandler) GetById(c *fiber.Ctx) error {
	log.Println("Entering BOD GetById....")
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	// if error in parsing string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
		})
	}
	data, err := bh.bodService.FindById(id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": data,
	})
}
