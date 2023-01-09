package config

import (
	"os"
	psg "rahmathidayat1203/backend-portofolio-v1/config/postgres"
	"strconv"

	"gorm.io/gorm"
)

type config struct{}

type Config interface {
	ServiceName() string
	ServicePort() int
	ServiceEnvironment() string
	Database() *gorm.DB
}

func NewConfig() Config {
	return &config{}
}
func (c *config) Database() *gorm.DB {
	return psg.InitDB()
}

func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}
func (c *config) ServicePort() int {
	v := os.Getenv("PORT")
	port, _ := strconv.Atoi(v)
	return port
}
func (c *config) ServiceEnvironment() string {
	return os.Getenv("ENV")
}
