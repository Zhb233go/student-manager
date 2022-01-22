package middleware

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

var Ln = logrus.New()

func Logs() func() {
	logName := time.Now().Format("2006-01-02 ") + ".log"
	file, err := os.OpenFile("./logs/"+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	Ln.SetReportCaller(true)
	Ln.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			"time":  "时间",
			"level": "级别",
			"msg":   "消息",
			"func":  "来源",
			"file":  "地址",
		},
	}
	Ln.SetOutput(io.MultiWriter(file, os.Stdout))
	return func() {
		if err := file.Close(); err != nil {
			Ln.Fatal(err)
		}
	}
}
