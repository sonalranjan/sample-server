package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sample-server/client"
	"github.com/sample-server/config"
	"github.com/sample-server/service"
)

type Handler struct {
	c *client.Client
	f *config.Config
	s *service.Service
}

func New(f *config.Config, c *client.Client, s *service.Service) *Handler {
	return &Handler{c, f, s}
}

func respond(w http.ResponseWriter, status int, data interface{}, err string) {
	response := &model.HTTPResponse{
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
