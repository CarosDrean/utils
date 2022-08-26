package bootstrap

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/CarosDrean/utils/model"
)

func newConfiguration() (model.Configuration, error) {
	// we ignore the error because it can take the system environments
	_ = godotenv.Load()

	portHttp, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return model.Configuration{}, err
	}

	return model.Configuration{
		PortHttp:       portHttp,
		AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowedMethods: strings.Split(os.Getenv("ALLOWED_METHODS"), ","),
	}, nil
}
