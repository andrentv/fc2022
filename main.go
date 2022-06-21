package main

import (
	route2 "github.com/codedede/imersaofsfc2.simulator/application/route"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	kafka.Publish("Ola", "readtest", producer)

	// route := route2.Route{
	// 	ID: "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[1])

}