package initialize

import (
	"DRCache/global"
	"DRCache/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitMysqlDB 初始化mysql数据库连接
func InitMysqlDB() {
	mysqlInfo := global.Settings.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.User, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Lg.Error("[InitMysqlDB] 连接mysql异常")
		global.Lg.Error(mysqlInfo.Host)
		global.Lg.Error(err.Error())
	}
	global.DB = db

	global.DB.AutoMigrate(&models.User{})
}

// InitPGDB 初始化pg数据库连接
func InitPGDB() {
	pgInfo := global.Settings.PGInfo
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			pgInfo.Host, pgInfo.User, pgInfo.Password, pgInfo.DBName, pgInfo.Port),
		PreferSimpleProtocol: true, // 禁用 prepared statement
	}), &gorm.Config{})
	if err != nil {
		global.Lg.Error("[InitPGDB] 连接pg异常")
		global.Lg.Error(pgInfo.Host)
		global.Lg.Error(err.Error())
	}
	global.DB = db

	//global.DB.AutoMigrate()
}
