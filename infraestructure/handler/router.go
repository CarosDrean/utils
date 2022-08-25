package handler

import (
	"utils/infraestructure/handler/ip"
	"utils/model"
)

func InitRoutes(specification model.RouterSpecification) {
	ip.NewRouter(specification)
}
