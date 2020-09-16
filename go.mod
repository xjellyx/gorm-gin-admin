module github.com/olongfen/gorm-gin-admin

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.2
	github.com/casbin/casbin/v2 v2.11.2
	github.com/dchest/captcha v0.0.0-20170622155422-6a29415a8364
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/jianfengye/collection v0.0.0-20200827021300-290144bad0d5 // indirect
	github.com/lib/pq v1.8.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/olongfen/contrib v0.0.0-20200916020620-d4cb00a9a2dd
	github.com/olongfen/gen-id v1.0.2-0.20200910083018-c2a28915b392
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.1
	github.com/subosito/gotenv v1.2.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	gorm.io/driver/postgres v1.0.0
	gorm.io/gorm v1.20.1

)

replace github.com/spf13/viper v1.7.1 => github.com/olongfen/viper v1.7.2
