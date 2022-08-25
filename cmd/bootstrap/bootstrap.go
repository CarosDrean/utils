package bootstrap

import (
	"fmt"

	"utils/infraestructure/handler"
	"utils/model"
)

func Run() error {
	config, err := newConfiguration()
	if err != nil {
		return fmt.Errorf("newConfiguration(): %v", err)
	}

	api := newEcho(config)

	handler.InitRoutes(model.RouterSpecification{
		Config: config,
		Api:    api,
	})

	port := fmt.Sprintf(":%d", config.PortHttp)
	return api.Start(port)
}
