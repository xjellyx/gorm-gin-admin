package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/olongfen/contrib/log"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
	"strings"
)

type Global struct {
	DB       Database
	RDB      RedisDB
	Serve    Serve
	FilePath FilePath
	Project  Project
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

type Serve struct {
	ServerAddr string
	ServerPort string
	RpcHost    string
	RpcPort    string
	IsTLS      bool // true: 开启https
	TLS        TLS
}

type FilePath struct {
	UserKeyDir   string // 私钥地址
	UserPubDir   string // 公钥地址
	AdminKeyDir  string // 私钥地址
	AdminPubDir  string // 公钥地址
	LogDir       string // 日志保存地址
	LogPatent    string // 日志格式
	HeadIconDir  string // 头像保存路径
	RBACModelDir string // casbin 保存路径
}

type Project struct {
	MaxRoleLevel     int
	LoginRestriction bool // 登录限制
}

var (
	Settings = new(Global)
	DevEnv   = false
)

func init() {
	var (
		err        error
		configFile string
	)
	if err = gotenv.Load("./conf/.env"); err != nil {
		log.Fatal(err)
	}
	env := os.Getenv("ENVIRONMENT")
	switch {
	case strings.Contains(env, "prod"):
		configFile = "./conf/prod-global-config.yaml"
	case strings.Contains(env, "test"):
		configFile = "./conf/test-global-config.yaml"
	default: // default is dev
		DevEnv = true
		configFile = "./conf/dev-global-config.yaml"

	}
	viper.SetConfigFile(configFile)
	viper.AddConfigPath(".")
	// database
	viper.SetDefault("db", Database{
		Host:         "0.0.0.0",
		Port:         "65433",
		Driver:       "postgres",
		DatabaseName: "business",
		Username:     "business",
		Password:     "business",
		MaxIdleConn:  200,
		MaxOpenConn:  2000,
	})
	// redis
	viper.SetDefault("rdb", RedisDB{
		Host:     "0.0.0.0",
		Port:     "6379",
		DB:       0,
		Password: "",
	})
	// serve
	viper.SetDefault("serve", Serve{
		ServerAddr: "0.0.0.0",
		ServerPort: "8050",
		RpcHost:    "0.0.0.0",
		RpcPort:    "9050",
		IsTLS:      false,
		TLS: TLS{
			Cert: "./conf/serve.cert",
			Key:  "./conf/serve.key",
		},
	})
	// file path
	viper.SetDefault("filepath", FilePath{
		UserKeyDir:   "./conf/user.key",
		UserPubDir:   "./conf/user.pub",
		AdminKeyDir:  "./conf/admin.key",
		AdminPubDir:  "./conf/admin.pub",
		LogDir:       "./log",
		LogPatent:    "_%Y-%m-%d.log",
		HeadIconDir:  "./public/static/",
		RBACModelDir: "./conf/model_casbin.conf",
	})
	viper.WatchConfig()
	// project
	viper.SetDefault("project", Project{MaxRoleLevel: 9})
	if err = viper.WriteConfig(); err != nil {
		log.Warnln(err)
	}
	_ = viper.ReadInConfig()
	if err = viper.Unmarshal(Settings); err != nil {
		log.Fatal(err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
		if err = viper.Unmarshal(Settings); err != nil {
			log.Fatal(err)
		}
	})
	log.Infoln("setting init success !")
}
