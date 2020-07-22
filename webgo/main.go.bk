package main

import (
    "fmt"
    //"os"

    "github.com/gin-contrib/static"
    "github.com/gin-gonic/gin"
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
        c.String(200, `{"message":"hello, hello, hello"}`)
    })

    if err := r.Run(fmt.Sprintf("%s:%d", args.BindAddress, args.BindPort)); err != nil {
        return err
    }

    return nil
}

func main() {
    args := arguments{
        LogLevel: "info",
        BindAddress: "0.0.0.0",
        BindPort: 9080,
        StaticContents: "./static",
        }

    if err := runServer(args); err != nil {
        logger.WithError(err).Fatal("Server exits with error")
    }
}
