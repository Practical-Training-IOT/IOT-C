package config

import "github.com/jackc/pgx/v5/pgxpool"

var (
	DB          *pgxpool.Pool
	AlarmStruct AlarmNacosStruct
	DataConfig  Configs
)
