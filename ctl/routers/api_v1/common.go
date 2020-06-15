package api_v1

import (
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/user_base/pkg/setting"
)

var (
	logCtl = log.NewLogFile("./log/log_ctrl", setting.ProjectSetting.IsProduct)
)
