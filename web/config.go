package web

import (
	"encoding/json"
	"mtworldview/app"
	"net/http"
)

//Public facing config
type PublicConfig struct {
	Version         string               `json:"version"`
}

type ConfigHandler struct {
	ctx *app.App
}

func (h *ConfigHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")

	webcfg := PublicConfig{}
	webcfg.Version = app.Version

	json.NewEncoder(resp).Encode(webcfg)
}
