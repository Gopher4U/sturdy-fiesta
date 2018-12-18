package handler

import "github.com/Gopher4U/sturdy-fiesta/RESTapi/config"

type ConfigHandler struct {
	*config.Config
}

func StartConfig (c *config.Config) *ConfigHandler{
	return &ConfigHandler{c}
}
