package config

import (
	"github.com/Kaivv1/blog-aggregator/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewConfig(db *database.Queries) *ApiConfig {
	return &ApiConfig{
		DB: db,
	}
}
