package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

type Msg struct {
	UserId    string   `json:"userId"`
	Text      string   `json:"text"`
	State     string   `json:"state"`
	Namespace string   `json:"namespace"`
	Rooms     []string `json:"rooms"`
}

func main() {
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		//msg := Msg{s.ID(), "connected!", "notice", "", nil}
		s.SetContext("")

		s.Emit("connect", "connect")
		fmt.Println("connected /:", s.ID())
		fmt.Println("connected /:", s.URL())
		// fmt.Printf("URL: %#v \n", s.URL())
		// fmt.Printf("LocalAddr: %#+v \n", s.LocalAddr())
		// fmt.Printf("RemoteAddr: %#+v \n", s.RemoteAddr())
		// fmt.Printf("RemoteHeader: %#+v \n", s.RemoteHeader())
		// fmt.Printf("Cookies: %s \n", s.RemoteHeader().Get("Cookie"))
		//s.Emit("connect")
		return nil
	})

	server.OnConnect("connection", func(s socketio.Conn) error {
		msg := Msg{s.ID(), "connected!", "notice", "", nil}
		s.SetContext("")
		s.Emit("res", msg)
		fmt.Println("connected /:", s.ID())
		// fmt.Printf("URL: %#v \n", s.URL())
		// fmt.Printf("LocalAddr: %#+v \n", s.LocalAddr())
		// fmt.Printf("RemoteAddr: %#+v \n", s.RemoteAddr())
		// fmt.Printf("RemoteHeader: %#+v \n", s.RemoteHeader())
		// fmt.Printf("Cookies: %s \n", s.RemoteHeader().Get("Cookie"))
		return nil
	})

	server.OnConnect("/connection", func(s socketio.Conn) error {
		msg := Msg{s.ID(), "connected!", "notice", "", nil}
		s.SetContext("")
		s.Emit("res", msg)
		fmt.Println("connected /:", s.ID())
		// fmt.Printf("URL: %#v \n", s.URL())
		// fmt.Printf("LocalAddr: %#+v \n", s.LocalAddr())
		// fmt.Printf("RemoteAddr: %#+v \n", s.RemoteAddr())
		// fmt.Printf("RemoteHeader: %#+v \n", s.RemoteHeader())
		// fmt.Printf("Cookies: %s \n", s.RemoteHeader().Get("Cookie"))
		return nil
	})
	server.OnEvent("/", "connection", func(s socketio.Conn, room string) {
		s.Join(room)
		//msg := Msg{s.ID(), "<= " + s.ID() + " join " + room, "state", s.Namespace(), s.Rooms()}
		fmt.Println("/:join", room, s.Namespace(), s.Rooms())
		//server.BroadcastToRoom(room, "res", msg)
	})

	server.OnEvent("/", "join", func(s socketio.Conn, room string) {
		s.Join(room)
		//msg := Msg{s.ID(), "<= " + s.ID() + " join " + room, "state", s.Namespace(), s.Rooms()}
		fmt.Println("/:join", room, s.Namespace(), s.Rooms())
		//server.BroadcastToRoom(room, "res", msg)
	})
	server.OnEvent("/", "leave", func(s socketio.Conn, room string) {
		s.Leave(room)
		//msg := Msg{s.ID(), "<= " + s.ID() + " leave " + room, "state", s.Namespace(), s.Rooms()}
		fmt.Println("/:chat received", room, s.Namespace(), s.Rooms())
		//server.BroadcastToRoom(room, "res", msg)
	})

	server.OnEvent("/", "chat", func(s socketio.Conn, msg string) {
		res := Msg{s.ID(), "<= " + msg, "normal", s.Namespace(), s.Rooms()}
		s.SetContext(res)
		//fmt.Println("/:chat received", msg, s.Namespace(), s.Rooms(), server.Rooms())
		rooms := s.Rooms()
		if len(rooms) > 0 {
			fmt.Println("broadcast to", rooms)
			//for i := range rooms {
			//server.BroadcastToRoom(rooms[i], "res", res)
			//}
		}
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("/:notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		fmt.Println("/chat:msg received", msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(Msg)
		s.Emit("bye", last)
		//res := Msg{s.ID(), "<= " + s.ID() + " leaved", "state", s.Namespace(), s.Rooms()}
		rooms := s.Rooms()
		s.LeaveAll()
		s.Close()
		if len(rooms) > 0 {
			fmt.Println("broadcast to", rooms)
			//for i := range rooms {
			//	server.BroadcastToRoom(rooms[i], "res", res)
			//}
		}
		fmt.Printf("/:bye last context: %#+v \n", s.Context())
		return last.Text
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("/:error ", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("/:closed", s.ID(), reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))

	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe("192.168.1.233:3000", nil))
}
