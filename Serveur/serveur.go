package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
) // only needed below for sample processing

// func non() {
// 	db := DbRedis()
// 	//Si vote Redis n'a pas la configuration basique utilisez :
// 	//db.NewRedis("addr","password",db)
// 	db.DefaultRedis()

// 	Get := func(rdb *redis.Client, key string) *redis.StringCmd {
// 		cmd := redis.NewStringCmd("get", key)
// 		rdb.Process(cmd)
// 		return cmd
// 	}

// }

func oui() {

	var book []string

	fmt.Println("Launching server...")
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")
	// accept connection on port
	conn, _ := ln.Accept()

	//ugly boucle
	for {
		// received messaged
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			book = append(book, string(message))
		} else if message == "STOP" {
			newmessage := strings.ToUpper("ok frr")
			conn.Write([]byte(newmessage + "\n"))
			break
		}
	}
}
