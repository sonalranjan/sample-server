package handler

import (
	"encoding/json"
	"net/http"

	"github.com/srnewbie/sample-server/client"
	"github.com/srnewbie/sample-server/config"
)

type (
	Handler struct {
		c *client.Client
		f *config.Config
	}
	HTTPResponse struct {
		Data interface{} `json:"data"`
		Err  string      `json:"err"`
	}
)

func New(f *config.Config, c *client.Client) *Handler {
	return &Handler{c, f}
}

func respond(w http.ResponseWriter, status int, data interface{}, err string) {
	response := &HTTPResponse{
		Data: data,
		Err:  err,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}

func (h *Handler) Health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		respond(w, 200, "ok", "")
	}
}
