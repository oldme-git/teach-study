package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wssrv *websocket.Conn
var conn map[int]int

func main() {
	http.HandleFunc("/", home)
	// 定义ws路径
	http.HandleFunc("/ws", ws)
	err := http.ListenAndServe("localhost:18000", nil)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func ws(w http.ResponseWriter, r *http.Request) {
	// 将http服务升级成ws服务
	wssrv, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer wssrv.Close()
	for {
		// 监听消息
		mt, message, err := wssrv.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		// 监听到信息，向客户端响应
		err = wssrv.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
