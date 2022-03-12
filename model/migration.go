package model

import (
	"registry-manager/pkg/conf"
	"registry-manager/pkg/util"
)

//执行数据迁移
func migration() {

	util.Log.Info().Msg("开始进行数据库初始化...")

	// 自动迁移模式
	if conf.DatabaseConfig.Type == "mysql" {
		DB = DB.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	DB.AutoMigrate(User{})
}
