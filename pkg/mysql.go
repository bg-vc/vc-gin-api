package pkg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"vc-gin-api/pkg/log"
)

type DBReadConfig struct {
	Schema   string
	Host     string
	Port     string
	Username string
	Password string
}

type DBWriteConfig struct {
	Schema   string
	Host     string
	Port     string
	Username string
	Password string
}

func LoadDBReadConfig(viper *viper.Viper) *DBReadConfig {
	cfg := &DBReadConfig{
		Schema:   viper.GetString("schema"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
	log.Infof("LoadDBReadConfig:%+v", cfg)
	return cfg
}

func LoadDBWriteConfig(viper *viper.Viper) *DBWriteConfig {
	cfg := &DBWriteConfig{
		Schema:   viper.GetString("schema"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
	log.Infof("LoadDBWriteConfig:%+v", cfg)
	return cfg
}

func (cfg *DBReadConfig) InitDBRead() *sqlx.DB {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	dbRead, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.Errorf(err, "InitDBRead error:%v", dataSourceName)
		panic("InitDBRead error")
	}
	return dbRead
}

func (cfg *DBWriteConfig) InitDBWrite() *sqlx.DB {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	dbWrite, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.Errorf(err, "InitDBWrite error:%v", dataSourceName)
		panic("InitDBWrite error")
	}
	return dbWrite
}
