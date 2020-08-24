package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/src/models"
	"github.com/olongfen/user_base/src/pkg/app"
	"github.com/olongfen/user_base/src/pkg/setting"
	"github.com/olongfen/user_base/src/utils"
	"strings"
)

// CheckUserAuth 验证用户
func CheckUserAuth(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		// 按空格分割
		tokenStr := ""
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) > 1 {
				tokenStr = parts[1]
			}
		} else {
			tokenStr = c.Request.Header.Get("token")
		}

		if tokenStr == "" {
			tokenStr = c.Query("token")
		}

		if tokenStr == "" {
			c.JSON(500, app.Response{
				Meta: app.Meta{
					Code:    500,
					Message: "token invalid",
				},
				Data: nil,
			})
			c.Abort()
			return
		}

		var (
			s   *session.Session
			err error
		)
		if isAdmin {
			if s, err = models.AdminKey.SessionDecode(tokenStr); err != nil {
				c.JSON(401, app.Response{
					Meta: app.Meta{
						Code:    401,
						Message: contrib.ErrTokenInvalid.Error(),
					},
					Data: nil,
				})
				c.Abort()
				return
			}
		} else {
			// 验证用户,管理员和普通用户的密钥对不一样，所以验证两次,管理员token可以使用与普通界面
			if s, err = models.UserKey.SessionDecode(tokenStr); err != nil {
				if s, err = models.AdminKey.SessionDecode(tokenStr); err != nil {
					c.JSON(401, app.Response{
						Meta: app.Meta{
							Code:    401,
							Message: contrib.ErrTokenInvalid.Error(),
						},
						Data: nil,
					})
					c.Abort()
					return
				}
			}
		}
		// 不是同一个ip地址
		if s.IP != c.ClientIP() && setting.Setting.IsProduct {
			c.JSON(403, app.Response{
				Meta: app.Meta{
					Code:    403,
					Message: utils.ErrIPAddressInvalid.Error(),
				},
				Data: nil,
			})
			c.Abort()
			return
		}
		c.Set("sessionTag", s)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
