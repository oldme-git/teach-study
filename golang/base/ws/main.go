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

var clients = make(map[*websocket.Conn]struct{})

func main() {
	http.HandleFunc("/", home)
	// 定义ws路径
	http.HandleFunc("/ws", ws)
	// 广播
	http.HandleFunc("/broadcast", broadcast)
	err := http.ListenAndServe("localhost:18000", nil)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func broadcast(w http.ResponseWriter, r *http.Request) {
	for client, _ := range clients {
		client.WriteMessage(websocket.TextMessage, []byte("能听到我的广播吗"))
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	// 将http服务升级成ws服务
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	// 建立连接将客户端添加到map中
	clients[c] = struct{}{}
	defer c.Close()
	for {
		// 监听消息
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		// 监听到信息，向客户端响应
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	// 关闭连接将客户端从map中剔除
	delete(clients, c)
}
