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
	for _, element := range toSave {
		word := strings.Fields(element)
		set := rdb.Do("SET", word[0], word[1])
		println(set)
	}
}

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
