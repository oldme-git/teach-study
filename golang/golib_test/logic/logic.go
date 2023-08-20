package logic

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"golib_test/consts"
	"os"
	"strconv"
	"sync"
)

var m sync.Mutex

func Reserve(ctx context.Context, uid string, sid int) {
	m.Lock()
	isUse, _ := getSeat(ctx, sid)
	if isUse {
		setSeat(ctx, sid)
		m.Unlock()
		saveSeat(ctx, uid, sid)
		g.Log().Infof(ctx, "uid:%s，sid:%d", uid, sid)
	} else {
		m.Unlock()
	}
}

func setSeat(ctx context.Context, sid int) (err error) {
	_, err = g.Redis().HSet(ctx, consts.SeatsKey, map[string]interface{}{
		strconv.Itoa(sid): false,
	})
	return
}

func getSeat(ctx context.Context, sid int) (bool, error) {
	res, err := g.Redis().HGet(ctx, consts.SeatsKey, strconv.Itoa(sid))
	if err != nil {
		return false, err
	}
	return res.Bool(), nil
}

func saveSeat(ctx context.Context, uid string, sid int) (err error) {
	hostname, _ := os.Hostname()
	_, err = g.Redis().HSet(ctx, consts.UseSeatsKey, map[string]interface{}{
		uid: strconv.Itoa(sid) + hostname,
	})
	return
}

func InitData() {
	var ctx = gctx.New()
	g.Redis().HSet(ctx, consts.SeatsKey, map[string]interface{}{
		"0": true,
		"1": true,
		"2": true,
	})
	g.Redis().Del(ctx, consts.UseSeatsKey)
}
