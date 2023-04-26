package goframe

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"testing"
)

// http和ws请求
func TestHttpWs(t *testing.T) {
	s := g.Server()

	// 用来储存ws对象
	var wsList []*ghttp.WebSocket

	s.BindHandler("/ws", func(r *ghttp.Request) {
		var ctx = r.Context()
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(ctx, err)
			r.Exit()
		}
		// 储存
		wsList = append(wsList, ws)

	})
	s.BindHandler("/", func(r *ghttp.Request) {
		// 发送信息
		for _, ws := range wsList {
			ws.WriteMessage(1, []byte("我是hello"))
		}
		r.Response.Write("哈喽世界！")
	})
	s.SetPort(8199)
	s.Run()
}
