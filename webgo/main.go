package main

import (
	"fmt"
	"strconv"
	//"os"
	"database/sql"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

var logLevelMap = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

type arguments struct {
	LogLevel       string
	BindAddress    string
	BindPort       int
	StaticContents string
}

type Tag struct {
	//Id   int    `json:"id"`
	Username string `json:"username"`
}

func runServer(args arguments) error {
	level, ok := logLevelMap[args.LogLevel]
	if !ok {
		return fmt.Errorf("Invalid log level: %s", args.LogLevel)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"args": args,
	}).Info("Given options")

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile(args.StaticContents, false)))
	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("Hello World1")
		db, err := sql.Open("mysql", "cedar:Degc2019@tcp(192.168.122.150:3306)/test")
		if err != nil {
			fmt.Println("you broke it")
			panic(err.Error())
		}
		fmt.Println("Hello World2")
		defer db.Close()
		//var tag Tag
		results, err := db.Query("SELECT Username FROM user")
		if err != nil {
			fmt.Println("error in part 2")
		}
		//err = results.Scan(&tag.Username)
		//fmt.Println(tag.Username, "Hello World 3")
		var tag Tag
		var content string
		for results.Next() {
			//var tag Tag
			// for each row, scan the result into our tag composite object
			err = results.Scan(&tag.Username)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			content = "{" + strconv.Quote("message") + ":" + strconv.Quote(tag.Username) + "}"

			//content = `{"message":"hello, hello, hello"}`
		}
		fmt.Println(content)
		c.String(200, content)
	})

	if err := r.Run(fmt.Sprintf("%s:%d", args.BindAddress, args.BindPort)); err != nil {
		return err
	}

	return nil
}

func main() {
	args := arguments{
		LogLevel:       "info",
		BindAddress:    "0.0.0.0",
		BindPort:       9080,
		StaticContents: "./static",
	}

	if err := runServer(args); err != nil {
		logger.WithError(err).Fatal("Server exits with error")
	}
}
