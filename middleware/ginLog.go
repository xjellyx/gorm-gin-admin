package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log2 "github.com/olongfen/contrib/log"
	"github.com/olongfen/user_base/pkg/setting"
	"time"
)

func GinAPILog() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			// 你的自定义格式
			return fmt.Sprintf(`address: %s, time: [%s], method: %s,  message: %s `+"\n"+`path: %s, proto: %s, codes: %d, latency: %s, agent: %s`+"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.ErrorMessage,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
			)
		},
		Output: log2.NewLogFile("./log/router", !setting.Setting.IsProduct, setting.Setting.LogPatent).Out,
	})

}
