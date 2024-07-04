package svc

import (
	"gorm.io/gorm"
	"notify/common"
	"notify/notify-api/internal/config"
	notify_model "notify/notify-model"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := common.InitGorm(c.MySQL.DataSource)
	mysqlDb.AutoMigrate(
		&notify_model.Notice{},
		&notify_model.User{},
	)
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
	}
}
