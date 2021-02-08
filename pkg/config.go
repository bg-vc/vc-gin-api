package pkg

import (
	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gopkg.in/redis.v5"
	"vc-gin-api/pkg/log"
)

var (
	Cfg      *Config
	DB       *sqlx.DB
	RedisCli *redis.Client
)

type Config struct {
	Logger *LoggerConfig
	DB     *DBConfig
	Redis  *RedisConfig
}

func Init(cfgName string) {
	setConfig(cfgName)
	Cfg = loadConfig()
	initConfig(Cfg)
	watchConfig()
}

func setConfig(cfgName string) {
	if cfgName != "" {
		viper.SetConfigFile(cfgName)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic("viper.ReadInConfig error")
	}
}

func loadConfig() *Config {
	cfg := &Config{
		Logger: LoadLoggerConfig(viper.Sub("logger")),
		DB:     LoadDbConfig(viper.Sub("db")),
		Redis:  LoadRedisConfig(viper.Sub("redis")),
	}

	return cfg
}

func initConfig(cfg *Config) {
	cfg.Logger.InitLogger()
	DB = cfg.DB.InitDB()
	//RedisCli = cfg.Redis.InitRedis()
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config file changed:%s", e.Name)
	})
}
