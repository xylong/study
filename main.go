package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)

	//logrus.Trace("trace msg")
	//logrus.Debug("debug msg")
	logrus.WithFields(logrus.Fields{
		"user_id": 10001,
		"ip":      "127.0.0.1",
	}).Info("hello world")
	//logrus.Warn("warn msg")
	//logrus.Error("error msg")
	//logrus.Fatal("fatal msg")
	//logrus.Panic("panic msg")
}
