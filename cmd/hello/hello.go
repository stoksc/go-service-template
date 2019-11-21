package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/stoksc/hello/internal/hello"
	"go.uber.org/zap"
)

const connStrTmpl = "host=%v port=%v user=%v dbname=%v password=%v %v"

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer l.Sync()

	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	dbConf := viper.GetStringMapString("db")
	connStr := fmt.Sprintf(connStrTmpl, dbConf["host"], dbConf["port"], dbConf["dbname"], dbConf["user"], dbConf["password"], dbConf["options"])
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Errorf("fatal error connecting to db: %s", err))
	}
	defer db.Close()

	helloService := hello.HelloService{db, l}

	r := gin.Default()

	r.GET("/greeting", helloService.GetBaseGreetingHandler)
	r.GET("/greeting/:name", helloService.GetGreetingHandler)
	r.POST("/greeting", helloService.CreateGreetingHandler)

	r.Run()
}
