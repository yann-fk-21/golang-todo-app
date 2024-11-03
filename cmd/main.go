package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/yann-fk-21/todo-app/cmd/api"
	"github.com/yann-fk-21/todo-app/config"
	"github.com/yann-fk-21/todo-app/db"
	"github.com/yann-fk-21/todo-app/logger"
)

func main() {

	logger := logger.InitLogger()
	cfgEnv := config.InitConfig() 
	
    cfg := mysql.Config {
		User: cfgEnv.User,
		Passwd: cfgEnv.Passwd,
		Addr: cfgEnv.Addr,
		DBName: cfgEnv.DBName,
		Net: cfgEnv.Net,
		AllowNativePasswords: true,
		ParseTime: true,
	}
	mysqlDB, err := db.NewMysqlStorage(cfg)
	if err != nil {
		logger.Println(err)
		logger.Fatal(err)
	}

	if err := db.InitStorage(mysqlDB); err != nil {
		logger.Println(err)
		logger.Fatal(err)
	}



	server := api.NewServer(cfgEnv.ServerPort, mysqlDB)
	if err := server.Run(); err != nil {
        logger.Println(err)
		logger.Fatal(err)
	}
}