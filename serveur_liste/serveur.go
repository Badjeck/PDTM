package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)


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
