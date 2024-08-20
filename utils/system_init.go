package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitConfig() {
	// 初始化配置文件
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read configuration file")
	}
	fmt.Println("app config", viper.Get("app"))
	fmt.Println("mysql config", viper.Get("mysql"))
}

var DB *gorm.DB

func InitMySQL() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 设置输出目标和日志前缀
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 是否彩色输出
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil, err
	}
	fmt.Println("Database connection successful!")
	return DB, nil
}

func GetDB() *gorm.DB {
	return DB
}
