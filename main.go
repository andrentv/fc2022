package main

import (
	// route2 "github.com/codedede/imersaofsfc2.simulator/application/route"
    	"github.com/codedede/imersaofsfc2.simulator/infra/kafka"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/codedede/imersaofsfc2.simulator/application/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)	
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("Ola", "readtest", producer)

	// for {
	// 	_ = 1
	// }

	// route := route2.Route{
	// 	ID: "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[1])

}
