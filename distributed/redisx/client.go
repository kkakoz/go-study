package redisx

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Client *redis.Client

var count = 1

func Out(v ...any) {
	fmt.Print(count, " ")
	fmt.Println(v...)
	count++
}

func init() {
	godotenv.Load("./redisx/.env")
	fmt.Println(os.Getenv("ADDR"))
	Client = redis.NewClient(&redis.Options{
		Network:   "",
		Addr:      os.Getenv("ADDR"),
		Dialer:    nil,
		OnConnect: nil,
		Password:  os.Getenv("PWD"),
		DB:        1,
	})
	err := Client.Ping().Err()
	if err != nil {
		log.Fatal("err = ", err)
	}
}
