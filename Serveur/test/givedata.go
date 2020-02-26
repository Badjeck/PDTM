package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/go-redis/redis"
)

type server struct {
	ip string
}

func (s *server) Send(text string) {
	conn, _ := net.Dial("tcp", s.ip)
	fmt.Fprint(conn, text)
}

func listen() []string {
	var book []string

	fmt.Println("Launching server...")
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")
	// accept connection on port

	var message string
	//ugly boucle
	for {
		conn, _ := ln.Accept()
		// received messaged
		message, _ = bufio.NewReader(conn).ReadString('\n')
		if message == "STOP" {
			fmt.Println("ok")
			break
		}
		if message != "" {
			fmt.Println(message)
			book = append(book, string(message))
		}
	}
	return book
}

//Get get un element dans la bdd
func Get(rdb *redis.Client, chuckNoList []string) []string {
	var returnList []string
	for _, element := range chuckNoList {
		cmd := redis.NewStringCmd("get", string(element))
		rdb.Process(cmd)
		data, err := cmd.Result()
		if err != nil {
			returnList = append(returnList, "MISS "+string(element))
		}
		returnList = append(returnList, string(element)+" "+data)
	}
	return returnList
}

func sendToClient(returnList []string) {
	client := server{ip: "192.168.1.11:8081"}
	for _, element := range returnList {
		fmt.Println(element)
		client.Send(element)
	}
	client.Send("STOP")
}

func main() {
	//nouveau Client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	link := listen()
	returnData := Get(rdb, link)
	fmt.Println(returnData)
	sendToClient(returnData)

}
