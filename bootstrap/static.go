package bootstrap

import (
	"embed"
	"encoding/json"
	"github.com/gin-contrib/static"
	"io/fs"
	"io/ioutil"
	"net/http"
	"registry-manager/pkg/conf"
)

const StaticFolder = "assets/build"

type staticVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// StaticFS 内置静态文件资源
var StaticFS static.ServeFileSystem

// StaticEmbed embed 静态文件资源
var StaticEmbed embed.FS

type GinEmbedFileSystem struct {
	http.FileSystem
}

// Exists 文件是否存在
func (e GinEmbedFileSystem) Exists(prefix string, path string) bool {
	if _, err := e.Open(path); err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string, index bool) static.ServeFileSystem {
	subFS, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return GinEmbedFileSystem{
		FileSystem: http.FS(subFS),
	}
}

// InitStatic 初始化静态资源文件
func InitStatic() {
	var err error
	StaticFS = EmbedFolder(StaticEmbed, StaticFolder, false)
	// 检查静态资源的版本
	f, err := StaticFS.Open("version.json")
	if err != nil {
		//util.Log().Warning("静态资源版本标识文件不存在，请重新构建或删除 statics 目录")
		return
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		//util.Log().Warning("无法读取静态资源文件版本，请重新构建或删除 statics 目录")
		return
	}

	var v staticVersion
	if err := json.Unmarshal(b, &v); err != nil {
		//util.Log().Warning("无法解析静态资源文件版本, %s", err)
		return
	}

	staticName := "registry-manager-frontend"

	if v.Name != staticName {
		//util.Log().Warning("静态资源版本不匹配，请重新构建或删除 statics 目录")
		return
	}

	if v.Version != conf.RequiredStaticVersion {
		//util.Log().Warning("静态资源版本不匹配 [当前 %s, 需要: %s]，请重新构建或删除 statics 目录", v.Version, conf.RequiredStaticVersion)
		return
	}

}
