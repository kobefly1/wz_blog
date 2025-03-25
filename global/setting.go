package global

import (
	"github.com/tour-book/blog-service/pkg/logger"
	"github.com/tour-book/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.SeverSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
)
