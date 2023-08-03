package goframe

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	logConfig := glog.DefaultConfig()
	logConfig.Flags = logConfig.Flags | glog.F_ASYNC | glog.F_FILE_SHORT
	logConfig.RotateSize = 2 * 1024 * 1024
	logConfig.RotateBackupLimit = 100
	logConfig.RotateBackupExpire = 100 * time.Hour * 24
	logConfig.Path = "./a"
	err := g.Log().SetConfig(logConfig)
	if err != nil {
		panic(err)
	}
	i := 0
	for {
		i++
		g.Log().Info(nil, "测试", i)
	}
}
