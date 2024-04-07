package config

import (
	"github.com/api-sekejap/internal/constant"
	"github.com/api-sekejap/internal/entity"
	"github.com/api-sekejap/pkg/database"
	"github.com/api-sekejap/pkg/redis"
)

type Config struct {
	App               `json:"app" yaml:"app"`
	database.Database `json:"database" yaml:"database"`
	redis.MemoryCache `json:"redis" yaml:"redis"`
	Service           `json:"service" yaml:"services"`
}

type App struct {
	Name        string       `json:"name" yaml:"name"`
	Port        string       `json:"port" yaml:"port"`
	Environment constant.Env `json:"env" yaml:"env"`
}

type Service struct {
	Storage storageService `json:"storage" yaml:"storage"`
	Auth    authService    `json:"oauth" yaml:"oauth"`
}

type storageService struct {
	Firebase entity.FirebaseStorage
}

type authService struct {
	Google   entity.GoogleAuth   `json:"google" yaml:"google"`
	Facebook entity.FacebookAuth `json:"facebook" yaml:"facebook"`
}

// To determine testing and actual environment, setup proper env to do a RnD and enable after-research-implementer.
// Development mode.
func (c *Config) IsDevelopmentMode() bool {
	return c.App.Environment == constant.EnvDevelopment
}

// Staging mode.
func (c *Config) IsStagingMode() bool {
	return c.App.Environment == constant.EnvStaging
}
