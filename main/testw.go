package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

func Copy(ws *websocket.Conn) {
	fmt.Printf("copyServer %#v\n", ws.Config())
	io.Copy(ws, ws) //websocket.Conn实现了Read()和Write()，所以可以直接Copy。此时会阻塞并不断复制返回接收到的内容
	fmt.Println("copyServer finished")
}
func ReadWrite(ws *websocket.Conn) {
	fmt.Printf("readWriteServer %#v\n", ws.Config())
	buf := make([]byte, 100)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf[:n])
		n, err = ws.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%q\n", buf[:n])
	}
	fmt.Println("readWriteServer finished")
}

// 当接收string时，websocket.Message.Receive(ws, &buf)和ws.Read(buf)功能类似。发送同理
func RecvSend(ws *websocket.Conn) {
	fmt.Printf("recvSendServer %#v\n", ws)
	for {
		var buf string
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf)
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%q\n", buf)
	}
	fmt.Println("recvSendServer finished")
}

// 当接收[]byte时，websocket.Message.Receive(ws, &buf)可以接收发送二进制文件
func RecvSendBinary(ws *websocket.Conn) {
	fmt.Printf("recvSendBinaryServer %#v\n", ws)
	for {
		var buf []byte
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%#v\n", buf)
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", buf)
	}
	fmt.Println("recvSendBinaryServer finished")
}

type T struct {
	Msg  string `json:"msg"`
	Path string `json:"path"`
}

func Json(ws *websocket.Conn) {
	fmt.Printf("jsonServer %#v\n", ws.Config())
	for {
		var msg T
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%#v\n", msg)
		err = websocket.JSON.Send(ws, msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", msg)
	}
	fmt.Println("jsonServer finished")
}

func web(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)

	if r.Method == "GET" { //如果请求方法为get显示login.html,并相应给前端
	} else {

		//否则走打印输出post接受的参数username和password

		fmt.Println(r.PostFormValue("user"))

		fmt.Println(r.PostFormValue("password"))

	}
}

func Echo(ws *websocket.Conn) {

	var err error

	for {

		var reply string

		//websocket接受信息

		if err = websocket.Message.Receive(ws, &reply); err != nil {

			fmt.Println("receive failed:", err)

			break

		}

		fmt.Println("reveived from client: " + reply)

		msg := "received:" + reply

		fmt.Println("send to client:" + msg)

		//这里是发送消息

		if err = websocket.Message.Send(ws, msg); err != nil {

			fmt.Println("send failed:", err)

			break

		}

	}

}

//func main() {
//	//接收websocket的路由地址
//	http.Handle("/copy", websocket.Handler(Copy))
//	http.Handle("/readWrite", websocket.Handler(ReadWrite))
//	http.Handle("/recvSend", websocket.Handler(RecvSend))
//	http.Handle("/recvSendBinary", websocket.Handler(RecvSendBinary))
//	//http.Handle("/socket.io/", websocket.Handler(Json))
//	//html页面
//	http.HandleFunc("/socket.io/", web)
//	//http.Handle("/socket.io/", websocket.Handler(Echo))
//
//	if err := http.ListenAndServe("192.168.1.233:3000", nil); err != nil {
//		log.Fatal("ListenAndServe:", err)
//	}
//
//}
