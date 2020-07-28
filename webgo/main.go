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
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Tag2 struct {
        Id        string    `json:"id"`
        Username  string `json:"username"`
        Password  string `json:"password"`
        Firstname string `json:"firstname"`
        Lastname  string `json:"lastname"`
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
        db, err := sql.Open("mysql", "cedar:Degc2019@tcp(192.168.122.150:3306)/test")
	if err != nil {
		fmt.Println("you broke it")
		panic(err.Error())
	}
		defer db.Close()
	r.GET("/api/v1/hello", func(c *gin.Context) {
		results, err := db.Query("SELECT * FROM user")
		if err != nil {
			fmt.Println("error in part 2")
		}
		var tag Tag
                var content string
                var temp string
		for results.Next() {
			err = results.Scan(&tag.Id, &tag.Username, &tag.Password, &tag.Firstname, &tag.Lastname)
			if err != nil {
				panic(err.Error())
			}
			temp = temp + strconv.Itoa(tag.Id)+tag.Username+tag.Password+tag.Firstname+tag.Lastname
                }
                content = "{" + strconv.Quote("message") + ":" + strconv.Quote(temp) + "}"
		c.String(200, content)
	})
	r.GET("/api/v1/add", func(c *gin.Context) {
		fmt.Println("hello")
		//insert, err := db.Query("insert into test values (1, 'Deathoath', '123456', 'Nami', 'Rakan')")
		insert, err := db.Query("insert into user values (1, 'Deathoath', '123456', 'Nami', 'Rakan')")
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		c.String(200, "{"+strconv.Quote("message")+":"+strconv.Quote("sent")+"}")
	})
        r.POST("api/v1/post", func(c *gin.Context){
                var tag Tag2
                var content string
		c.Bind(&tag)
                fmt.Println(tag.Id)
                content = "insert into user values(" + tag.Id + "," +strconv.Quote(tag.Username) + "," +strconv.Quote(tag.Password) + "," +strconv.Quote(tag.Firstname) + "," +strconv.Quote(tag.Lastname)+")"
                fmt.Println(content)
                insert, err := db.Query(content)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
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
