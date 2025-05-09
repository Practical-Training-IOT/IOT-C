package config

import (
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	AlarmStruct AlarmNacosStruct
	DataConfig  Configs
)
