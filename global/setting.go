package global

import (
	"github.com/tour-book/blog-service/pkg/logger"
	"github.com/tour-book/blog-service/pkg/setting"
)

// 配置结构体的全局变量
var (
	ServerSetting   *setting.SeverSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
)
