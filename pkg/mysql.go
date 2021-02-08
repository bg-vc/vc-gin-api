package pkg

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"vc-gin-api/pkg/log"
)

type DBConfig struct {
	Schema   string
	Host     string
	Port     string
	Username string
	Password string
}

func LoadDbConfig(viper *viper.Viper) *DBConfig {
	cfg := &DBConfig{
		Schema:   viper.GetString("schema"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
	fmt.Printf("LoadDbConfig:%+v\n", cfg)
	return cfg
}

func (cfg *DBConfig) InitDB() *sqlx.DB {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.Errorf(err, "InitDb error:%v", dataSourceName)
		panic("InitDb error")
	}
	return db
}
