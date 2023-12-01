package handler

import (
	"encoding/json"
	"net/http"

	"github.com/srnewbie/sample-server/client"
	"github.com/srnewbie/sample-server/config"
	"github.com/srnewbie/sample-server/service"
)

type (
	Handler struct {
		c *client.Client
		f *config.Config
		s *service.Service
	}
	HTTPResponse struct {
		Data interface{} `json:"data"`
		Err  string      `json:"err"`
	}
)

func New(f *config.Config, c *client.Client, s *service.Service) *Handler {
	return &Handler{
		c: c,
		f: f,
		s: s,
	}
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
		/*business logic*/
		h.s.HealthService()
		respond(w, 200, "ok", "")
	}
}
