package api

import (
	"net/http"

	"github.com/SENERGY-Platform/kong_admin/internal/config"
	"github.com/SENERGY-Platform/kong_admin/internal/kong_client"
)

type RouteHandler struct {
	Config     *config.Config
	KongClient *kong_client.KongClient
}

func (h *RouteHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	routes, err := h.KongClient.LoadRoutes()
	if err != nil {
		return err
	}
	SendResponse(w, routes)
	return nil
}

func New(kongClient *kong_client.KongClient, config *config.Config) HTTPHandler {
	return HTTPHandler{
		Handler: &RouteHandler{Config: config, KongClient: kongClient},
		Methods: []string{"GET"},
	}
}
