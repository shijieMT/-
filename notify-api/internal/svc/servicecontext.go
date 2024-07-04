package svc

import (
	"gorm.io/gorm"
	"notify/common"
	"notify/notify-api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := common.InitGorm(c.MySQL.DataSource)
	mysqlDb.AutoMigrate()
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
	}
}
