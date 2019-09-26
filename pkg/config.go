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
	DBRead   *sqlx.DB
	DBWrite  *sqlx.DB
	RedisCli *redis.Client
)

type Config struct {
	Logger  *LoggerConfig
	DBRead  *DBReadConfig
	DBWrite *DBWriteConfig
	Redis   *RedisConfig
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
		Logger:  LoadLoggerConfig(viper.Sub("logger")),
		DBRead:  LoadDBReadConfig(viper.Sub("db_read")),
		DBWrite: LoadDBWriteConfig(viper.Sub("db_write")),
		Redis:   LoadRedisConfig(viper.Sub("redis")),
	}

	return cfg
}

func initConfig(cfg *Config) {
	cfg.Logger.InitLogger()
	DBRead = cfg.DBRead.InitDBRead()
	DBWrite = cfg.DBWrite.InitDBWrite()
	RedisCli = cfg.Redis.InitRedis()
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config file changed:%s", e.Name)
	})
}
