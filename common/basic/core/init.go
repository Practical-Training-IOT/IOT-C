package core

import (
	"encoding/json"
	"fmt"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

var (
	err error
)

func init() {
	ZapInit()
	ConfigInit()
	NacosInit()
	PostGreSQLInit()
}

func ZapInit() {
	con := zap.NewDevelopmentConfig()
	con.OutputPaths = []string{
		"./zap.log",
		"stdout",
	}
	build, _ := con.Build()
	zap.ReplaceGlobals(build)
}

func ConfigInit() {
	viper.SetConfigFile("./config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		zap.S().Error(err)
		return
	}
	err = viper.Unmarshal(&config.DataConfig)
	if err != nil {
		zap.S().Error(err)
		return
	}
	fmt.Println(config.DataConfig)
}

func NacosInit() {
	C := config.DataConfig
	os.MkdirAll("./tmp/nacos/log", 0777)
	os.MkdirAll("./tmp/nacos/cache", 0777)
	clientConfig := constant.ClientConfig{
		NamespaceId:         C.NameSpace, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      C.Host,
			ContextPath: "/nacos",
			Port:        C.Port,
			Scheme:      "http",
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		zap.S().Error(err)
		return
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: C.DataId,
		Group:  C.Group})
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: C.DataId,
		Group:  C.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	err = json.Unmarshal([]byte(content), &config.AlarmStruct)
	if err != nil {
		zap.S().Error(err)
		return
	}
}

func PostGreSQLInit() {
	C := config.AlarmStruct.PostgreSQL
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", C.User, C.Password, C.Host, C.Port, C.Datebase)
	fmt.Println(dsn)
	config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Errorf("failed to connect database: %v", err)
		return
	}

	// 获取底层sql.DB对象并配置连接池
	sqlDB, err := config.DB.DB()
	if err != nil {
		zap.S().Errorf("failed to get sql.DB: %v", err)
		return
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可重用的最长时间

	zap.S().Info("Successfully connected to database with GORM")
}
