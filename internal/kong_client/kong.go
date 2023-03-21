package kong_client

import (
	"github.com/SENERGY-Platform/kong_admin/internal/config"
)

type KongClient struct {
	KONG_ADMIN_URL string
}

func LoadKongClient(config *config.Config) *KongClient {
	return &KongClient{KONG_ADMIN_URL: config.KONG_ADMIN_URL}
}
