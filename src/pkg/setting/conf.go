package setting

import (
	"github.com/olongfen/contrib/config"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"github.com/sirupsen/logrus"
	"time"
)

// ProjectConfig
type ProjectConfig struct {
	config.Config `yaml:"-"`
	ServerAddr    string
	ServerPort    string
	RpcHost       string
	RpcPort       string
	Sync          bool   // true: 数据库同步
	IsProduct     bool   //
	UserKeyDir    string // 私钥地址
	UserPubDir    string // 公钥地址
	AdminKeyDir   string // 私钥地址
	AdminPubDir   string // 公钥地址
	Db            *Database
	RDB           *RedisDB
	IsTLS         bool // true: 开启https
	TLS           *TLS
	IsCaptcha     bool
	LogDir        string
	LogPatent     string
	HeadIconDir   string
	RBACModelDir  string
	// RABCPolicyDir string
}

// RedisDB 缓存的连接参数
type RedisDB struct {
	Host     string
	Port     string
	DB       int
	Password string
}

// TLS
type TLS struct {
	Cert string `json:"cert"` // 证书文件
	Key  string `json:"key"`  // Key 文件
}

// Database 数据库连接参数
type Database struct {
	Host         string
	Port         string
	Driver       string
	DatabaseName string
	Username     string
	Password     string
	MaxIdleConn  int
	MaxOpenConn  int
}

var (
	Setting = &ProjectConfig{
		ServerAddr: utils.PubGetEnvString("SERVER_ADDR", "127.0.0.1"),
		ServerPort: utils.PubGetEnvString("SERVER_PORT", "8050"),
		RpcHost:    utils.PubGetEnvString("RPC_HOST", "127.0.0.1"),
		RpcPort:    utils.PubGetEnvString("RPC_PORT", "9050"),
		Sync:       false,
		IsProduct:  false,
		IsTLS:      false,
		Db: &Database{
			Driver:       utils.PubGetEnvString("DB_DRIVER", "postgres"),
			Host:         utils.PubGetEnvString("DB_HOST", "127.0.0.1"),
			Port:         utils.PubGetEnvString("DB_PORT", "65432"),
			DatabaseName: utils.PubGetEnvString("DB_NAME", "business"),
			Username:     utils.PubGetEnvString("DB_USERNAME", "business"),
			Password:     utils.PubGetEnvString("DB_PASSWORD", "business"),
			MaxIdleConn:  200,
			MaxOpenConn:  2000,
		},
		RDB: &RedisDB{
			Host:     utils.PubGetEnvString("RDB_HOST", "127.0.0.1"),
			Port:     utils.PubGetEnvString("RDB_PORT", "6379"),
			DB:       0,
			Password: utils.PubGetEnvString("RDB_PASSWORD", ""),
		},
		UserKeyDir:   "./conf/user.key",
		UserPubDir:   "./conf/user.pub",
		AdminKeyDir:  "./conf/admin.key",
		AdminPubDir:  "./conf/admin.pub",
		LogDir:       "./log",
		HeadIconDir:  "./public/static/",
		RBACModelDir: "./conf/model_casbin.conf",
		LogPatent:    "_%Y-%m-%d.log",
		// RABCPolicyDir: "./conf/policy_api.csv",
	}
)

// InitConfig 初始化配置文件
func InitConfig() {
	var (
		err error
	)
	if err = config.LoadConfigAndSave("./conf/project.config.yaml", Setting, Setting, time.Second*10); err != nil {
		logrus.Fatal(err)
	}
	if Setting.IsTLS {
		Setting.TLS = &TLS{
			Cert: "./conf/serve.cert",
			Key:  "./conf/serve.key",
		}
	}
	if err = Setting.Save(nil); err != nil {
		logrus.Fatal(err)
	}

}
