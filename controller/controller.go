package controller

import (
	"log"
	"sygap-cmdb/usecase"
)

type Icontroller interface {
}

type controller struct {
	model  usecase.Imodel
	logger *log.Logger
}

func NewController(m usecase.Imodel, logger *log.Logger) *controller {
	return &controller{m, logger}
}
