package model

type Configuration struct {
	PortHttp       int      `json:"port_http"`
	AllowedOrigins []string `json:"allowed_origins"`
	AllowedMethods []string `json:"allowed_methods"`
}
