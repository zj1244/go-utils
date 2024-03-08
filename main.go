package main

import (
	"github.com/zj1244/go-utils/log"
	"github.com/zj1244/go-utils/tar"
	"time"
)

func main() {
	var x string
	start := time.Now()
	tar.Untar("d:\\base.tar.gz", "d:\\tmp")
	x = time.Since(start).String()
	log.Infof("运行时间:%s", x)
}
