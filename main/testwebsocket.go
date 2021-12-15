package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader0 = websocket.Upgrader{
		//允许跨域访问
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler0(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello"))
	//收到http请求(upgrade),完成websocket协议转换
	//在应答的header中放上upgrade:websoket
	var (
		conn *websocket.Conn
		err  error
		//msgType int
		data []byte
	)
	if conn, err = upgrader0.Upgrade(w, r, nil); err != nil {
		//报错了，直接返回底层的websocket链接就会终断掉
		return
	}
	//得到了websocket.Conn长连接的对象，实现数据的收发
	for {
		//Text(json), Binary
		//if _, data, err = conn.ReadMessage(); err != nil {
		if _, data, err = conn.ReadMessage(); err != nil {
			//报错关闭websocket
			goto ERR
		}
		//发送数据，判断返回值是否报错
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			//报错了
			goto ERR
		}
	}
	//error的标签
ERR:
	conn.Close()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("connect"))
}

//func main() {
//	//http://localhost:7777/ws
//	http.HandleFunc("/ws", wsHandler0)
//	http.HandleFunc("/", helloHandler)
//	//http.HandleFunc("/", helloHandler)
//	//http.HandleFunc("/websocket", helloHandler)
//	//服务端启动
//	http.ListenAndServe("192.168.1.233:3000", nil)
//}
