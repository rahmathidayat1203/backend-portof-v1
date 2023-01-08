package config

import (
	psg "rahmathidayat1203/backend-portofolio-v1/config/postgres"

	"gorm.io/gorm"
)

type config struct{}

type Config interface {
	Database() *gorm.DB
}

func NewConfig() Config {
	return &config{}
}
func (c *config) Database() *gorm.DB {
	return psg.InitDB()
}
