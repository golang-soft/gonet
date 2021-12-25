package login

import (
	"context"
	"encoding/json"
	"fmt"
	"gonet/rpc"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

func GrpcAddEquip(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewGreeterClient(conn)
	if c == nil {
		log.Fatalf("链接失败")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ss, err := c.AddEquip(ctx, &rpc.AddEquipData{Uuid: "lzy111", ItemIdx: []int32{2001, 2002, 2003}})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	log.Printf("####### get server Greeting response: %s", ss.Data)

	data, err := json.Marshal(ss)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%v", data)
	// 返回json字符串给客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ss)
}

func GrpcAddHero(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewGreeterClient(conn)
	if c == nil {
		log.Fatalf("链接失败")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ss, err := c.AddHero(ctx, &rpc.AddHeroData{Uuid: "lzy111", HeroType: []int32{1, 2}})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	log.Printf("####### get server Greeting response: %s", ss.Data)

	data, err := json.Marshal(ss)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%v", data)
	// 返回json字符串给客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ss)
}

func GrpcAddItem(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewGreeterClient(conn)
	if c == nil {
		log.Fatalf("链接失败")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	items := []*rpc.Item{&rpc.Item{Id: 1, Count: 1, Type: 1}}
	ss, err := c.AddItem(ctx, &rpc.AddItemData{Uuid: "lzy111", ItemList: items})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	log.Printf("####### get server Greeting ressaponse: %s", ss.Data)

	data, err := json.Marshal(ss)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%v", data)
	// 返回json字符串给客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ss)
}

func GrpcGetRooms(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewGreeterClient(conn)
	if c == nil {
		log.Fatalf("链接失败")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ss, err := c.GetRooms(ctx, &rpc.ReqGetRooms{Start: 1, Len: 2})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	log.Printf("####### get server Greeting ressaponse: %s", ss.Data)

	data, err := json.Marshal(ss)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%v", data)
	// 返回json字符串给客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ss)
}
