package main

import (
	"github.com/sirupsen/logrus"
	"github.com/zj1244/go-utils/log"
	"github.com/zj1244/go-utils/tar"
	"time"
)

func setSaveDir(saveDir string) func(*log.LogrusConfig) {
	return func(logrusConfig *log.LogrusConfig) {
		logrusConfig.FileLocation = saveDir
	}
}
func init() {
	log.InitLogging(logrus.InfoLevel)
	log.InitLogging(logrus.InfoLevel, setSaveDir("test.log"))
}
func main() {
	var x string
	start := time.Now()
	tar.Untar("d:\\ab.tar", "d:\\tmp")
	x = time.Since(start).String()
	log.Infof("运行时间:%s", x)
}
