package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"registry-manager/pkg/conf"
	"registry-manager/pkg/util"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

// Init 初始化 MySQL 链接
func Init() {
	util.Log.Info().Msg("初始化数据库连接")

	var (
		db  *gorm.DB
		err error
	)

	if gin.Mode() == gin.TestMode {
		// 测试模式下，使用内存数据库
		db, err = gorm.Open(sqlite.Open("sqlite3"))
	} else {
		switch conf.DatabaseConfig.Type {
		case "UNSET", "sqlite", "sqlite3":
			// 未指定数据库或者明确指定为 sqlite 时，使用 SQLite3 数据库
			db, err = gorm.Open(sqlite.Open(util.RelativePath(conf.DatabaseConfig.DBFile)), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   conf.DatabaseConfig.TablePrefix,
					SingularTable: true,
				},
			})
		case "postgres":
			db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Name,
				conf.DatabaseConfig.Port)), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   conf.DatabaseConfig.TablePrefix,
					SingularTable: true,
				},
			})
		case "mysql", "mssql":
			db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.Port,
				conf.DatabaseConfig.Name,
				conf.DatabaseConfig.Charset)), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   conf.DatabaseConfig.TablePrefix,
					SingularTable: true,
				},
			})
		default:
			util.Log.Panic().Msgf("不支持数据库类型: %s", conf.DatabaseConfig.Type)
		}
	}

	//db.SetLogger(util.Log())
	if err != nil {
		util.Log.Panic().Err(err).Msg("连接数据库不成功")
	}

	// Debug模式下，输出所有 SQL 日志
	sqlDB, err := db.DB()
	if err != nil {
		util.Log.Panic().Err(err).Msg("连接数据库不成功")
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(50)
	//打开
	sqlDB.SetMaxOpenConns(100)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	DB = db

	//执行迁移
	migration()
}
