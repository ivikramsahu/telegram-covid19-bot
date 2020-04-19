package dataredis

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-redis/redis"
)

//GetDataSource should return value that has key :: value
func GetDataSource(sourceName string) string {

	defer EmptySet()
  if sourceName == "/country"{
    return EmptySet()
  }

	var validSource = regexp.MustCompile(`(?m)^\/(country)|(state)@([a-z]*)`)

	var globalSource = regexp.MustCompile(`(?m)\/globalstatus@global`)

	if validSource.MatchString(sourceName) || globalSource.MatchString(sourceName) {
		splitCountry := strings.Split(sourceName, "@")

		if splitCountry[1] == "global" {
			splitCountry[1] = strings.Replace(splitCountry[0], "/", "", 2)
		}

		sourceName = strings.Replace(splitCountry[1], "_", "-", 5)

		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		datasize := make(map[string]string)
		datasize, _ = client.HGetAll(sourceName).Result()

		fmt.Println(datasize)
		bodyData := new(bytes.Buffer)

		for key, value := range datasize {
			fmt.Fprintf(bodyData, "%s  → \"%s\"\n", key, value)
		}

		return bodyData.String()
	}

	return EmptySet()
}

//EmptySet default return
func EmptySet() string {

	return "No Data found... \nPlease Read /aboutbot"

}
