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
	if err != nil {
		log.Println(err)
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Search All", listBod)
}

func (bh *BodHandler) GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return response.HandleErrors(c, err)
	}
	data, _ := bh.bodService.FindById(id)
	return response.NewSuccess(c, fiber.StatusOK, "Search By Id", data)
}
