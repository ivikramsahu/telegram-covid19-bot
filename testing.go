package main

import (
  "github.com/go-redis/redis"
  "fmt"
)

func main() {
  client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  datasize := make(map[string]string)
  datasize, _ = client.HGetAll("india").Result()
  fmt.Println(datasize["NewConfirmed"])
  pong, err := client.Ping().Result()
  fmt.Println(pong, err)
 	 // Output: PONG <nil>
}



