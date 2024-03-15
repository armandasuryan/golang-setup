package controller

import (
	"log"
	"sygap-cmdb/usecase"

	"github.com/gofiber/fiber/v2"
)

type Icontroller interface {
	GetCiClassList(ctx *fiber.Ctx) error
}

type controller struct {
	model  usecase.Imodel
	logger *log.Logger
}

func NewController(m usecase.Imodel, logger *log.Logger) *controller {
	return &controller{m, logger}
}
