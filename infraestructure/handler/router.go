package handler

import (
	"github.com/CarosDrean/utils/infraestructure/handler/ip"
	"github.com/CarosDrean/utils/model"
)

func InitRoutes(specification model.RouterSpecification) {
	ip.NewRouter(specification)
}
