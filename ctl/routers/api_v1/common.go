package api_v1

import (
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/userDetail/pkg/setting"
)

var(
	logCtl=log.NewLogFile("./log/log_ctrl",setting.ProjectSetting.IsProduct)
)
