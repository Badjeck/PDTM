package main

import (
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
	for _,element := range returnList{
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

	test := []string{"l1_s1", "l1_s2","l1_s3","l1_s4"}
	returnData := Get(rdb, test)
	fmt.Println(returnData)
	sendToClient(returnData)

}
