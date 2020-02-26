package main

import (
	"bufio"
	"fmt"
	"net"
	"sort"
	"strings"
)

var res []string

type server struct {
	ip string
}

func (s *server) Send(text string) {
	fmt.Println(text, s.ip)
	conn, _ := net.Dial("tcp", s.ip)

	fmt.Fprint(conn, text)
}

func listen() []string {
	var book []string
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	var message string

	// elle est vraiment moche celle la
	for {
		conn, _ := ln.Accept()
		// recois un message
		message, _ = bufio.NewReader(conn).ReadString('\n')

		if message == "STOP" {
			break
		}
		if message != "" {
			book = append(book, string(message))
		}
	}
	return book
}

func getData(link []string) {
	s1 := server{ip: "192.168.242.2:8081"}
	s2 := server{ip: "192.168.242.3:8081"}
	s3 := server{ip: "192.168.242.4:8081"}
	s4 := server{ip: "192.168.242.5:8081"}
	for _, element := range link {
		link := strings.Fields(element)
		switch link[0] {
		case "l1_s1":
			//remove first alement of Slice
			oui := append(link[:0:0], link...)
			copy(oui[0:], oui[0+1:])
			oui[len(oui)-1] = ""
			oui = oui[:len(oui)-1]
			//send link to server
			for _, el := range oui {
				s1.Send(el)
				res1 := listen()
				res = addTab(res1, res)
			}
		case "l1_s2":
			//remove first alement of Slice
			oui := append(link[:0:0], link...)
			copy(oui[0:], oui[0+1:])
			oui[len(oui)-1] = ""
			oui = oui[:len(oui)-1]
			//send link to server
			for _, el := range oui {
				s2.Send(el)
				res1 := listen()
				res = addTab(res1, res)
			}
		case "l1_s3":
			//remove first alement of Slice
			oui := append(link[:0:0], link...)
			copy(oui[0:], oui[0+1:])
			oui[len(oui)-1] = ""
			oui = oui[:len(oui)-1]
			//send link to server
			for _, el := range oui {
				s3.Send(el)
				res1 := listen()
				res = addTab(res1, res)
			}
		case "l1_s4":
			//remove first alement of Slice
			oui := append(link[:0:0], link...)
			copy(oui[0:], oui[0+1:])
			oui[len(oui)-1] = ""
			oui = oui[:len(oui)-1]
			//send link to server
			for _, el := range oui {
				s4.Send(el)
				res1 := listen()
				res = addTab(res1, res)
			}
		}
	}
}

func addTab(tabsrc []string, tabcible []string) []string {
	for _, text := range tabsrc {
		tabcible = append(tabcible, text)
	}
	return tabcible
}

func structData() string {
	var data string

	sort.Strings(res)
	for _, linky := range res {
		strtab := strings.Fields(linky)
		for _, val := range strtab[1:] {
			data = data + val
		}
	}
	return data
}

func main() {

	dataLink := listen()
	getData(dataLink)
	oui := structData()
	fmt.Println(oui)
}
