package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/go-redis/redis"
)

// "github.com/go-redis/redis"

//DbRedis is your dataBase
type DbRedis struct {
	addr     string
	password string
	DB       int
}

//RedSave save a array in Redis
func RedSave(rdb *redis.Client, toSave []string) {
	var data11 string
	var data12 string
	var data13 string
	var data14 string

	var data21 string
	var data22 string
	var data23 string
	var data24 string
	for _, element := range toSave {
		word := strings.Fields(element)
		switch word[0]{
		case "l1_s1":
			data11 = data11 + " " + word[1]
		case "l1_s2":
			data12 = data12 + " " + word[1]
		case "l1_s3":
			data13 = data13 + " " + word[1]
		case "l1_s4":
			data14 = data14 + " " + word[1]
		case "l2_s1":
			data21 = data21 + " " + word[1]
		case "l2_s2":
			data22 = data22 + " " + word[1]
		case "l2_s3":
			data23 = data23 + " " + word[1]
		case "l2_s4":
			data24 = data24 + " " + word[1]
		}
	}
	rdb.Do("SET","l1_s1", data11)
	rdb.Do("SET","l1_s1", data12)
	rdb.Do("SET","l1_s1", data13)
	rdb.Do("SET","l1_s1", data14)

	rdb.Do("SET","l1_s1", data21)
	rdb.Do("SET","l1_s1", data22)
	rdb.Do("SET","l1_s1", data23)
	rdb.Do("SET","l1_s1", data24)

}

//Get ta bdd
func Get(rdb *redis.Client, key string) *redis.StringCmd{
		cmd := redis.NewStringCmd("get", key)
		rdb.Process(cmd)
		return cmd
	}

// }

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


func main() {
	db := DbRedis{addr: "localhost:6379", password: "", DB: 0}

	rdb := redis.NewClient(&redis.Options{
		Addr:     db.addr,
		Password: db.password,
		DB:       db.DB,
	})

	ListToSave := listen()
	fmt.Println(ListToSave)
	RedSave(rdb, ListToSave)
}
