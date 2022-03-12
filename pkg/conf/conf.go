package conf

import (
	"github.com/go-playground/validator/v10"
	"gopkg.in/ini.v1"
	"registry-manager/pkg/util"
)

// database 数据库
type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	DBFile      string
	Port        int
	Charset     string
}

// system 系统通用配置
type system struct {
	Listen        string `validate:"required"`
	Debug         bool
	SessionSecret string
	HashIDSalt    string
}

// DatabaseConfig 数据库配置
var DatabaseConfig = &database{
	Type:    "UNSET",
	Charset: "utf8",
	DBFile:  "registry-manager.db",
	Port:    3306,
}

// SystemConfig 系统公用配置
var SystemConfig = &system{
	Debug:  false,
	Listen: ":5212",
}

var cfg *ini.File

const defaultConf = `[System]
Listen = :5212
SessionSecret = {SessionSecret}
HashIDSalt = {HashIDSalt}
`

// Init 初始化配置文件
func Init(path string) {
	var err error
	if path == "" || !util.Exists(path) {
		// 创建初始配置文件
		confContent := util.Replace(map[string]string{
			"{SessionSecret}": util.RandStringRunes(64),
			"{HashIDSalt}":    util.RandStringRunes(64),
		}, defaultConf)
		f, err := util.CreatNestedFile(path)
		if err != nil {
			util.Log.Panic().Msgf("无法创建配置文件, %s", err)
		}

		// 写入配置文件
		_, err = f.WriteString(confContent)
		if err != nil {
			util.Log.Panic().Msgf("无法写入配置文件, %s", err)
		}

		f.Close()
	}

	cfg, err = ini.Load(path)
	if err != nil {
		util.Log.Panic().Msgf("无法解析配置文件 '%s': %s", path, err)
	}

	sections := map[string]interface{}{
		"Database": DatabaseConfig,
		"System":   SystemConfig,
	}
	for sectionName, sectionStruct := range sections {
		err = mapSection(sectionName, sectionStruct)
		if err != nil {
			util.Log.Panic().Msgf("配置文件 %s 分区解析失败: %s", sectionName, err)
		}
	}

	// 重设log等级
	if !SystemConfig.Debug {
		//util.Level = util.LevelInformational
		//util.GloablLogger = nil
		//util.Log()
	}

}

// mapSection 将配置文件的 Section 映射到结构体上
func mapSection(section string, confStruct interface{}) error {
	err := cfg.Section(section).MapTo(confStruct)
	if err != nil {
		return err
	}

	// 验证合法性
	validate := validator.New()
	err = validate.Struct(confStruct)
	if err != nil {
		return err
	}

	return nil
}
