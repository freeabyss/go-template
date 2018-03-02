package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"database/sql"
	"maizuo.com/back-end/go-template/src/server/common"
	"time"
)

func SetupDB()  {

	dialect := viper.GetString("db.dialect")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	database := viper.GetString("db.database")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	maxIdle := viper.GetInt("db.maxIdle")
	maxOpen := viper.GetInt("db.maxOpen")
	connMaxLifetime := viper.GetInt("db.connMaxLifetime")

	url := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open(dialect, url)
	if err != nil {
		panic("failed to connect database")
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime)*time.Millisecond)

	connectErr := db.Ping()      //测试连接
	if connectErr != nil {
		panic("failed to connect database:" + connectErr.Error())
	}

	common.DB = db
}

