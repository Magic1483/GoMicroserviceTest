package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	// "time"

	"github.com/segmentio/kafka-go"
)

type Message struct {
	Id 			int		`json:"Id"`
	Text 		string	`json:"Text"`
	IsChecked	bool	`json:"IsChecked"`
}


const (
	SendTopic = "handler_to_server"
	RecieveTopic = "server_to_handler"
	Address = "kafka:9092"

)



func ProduceMsg(ctx context.Context,msg []byte,key string) {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{Address},
		Topic: SendTopic,
	})

	err := w.WriteMessages(ctx,kafka.Message{
		Key: []byte(key),
		Value: []byte(msg),
	})
	if err != nil {
		panic("can't write msg "+err.Error())
	}
}

func consume (ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{Address},
		Topic: RecieveTopic,
		MinBytes: 5,
		MaxBytes: 1e6,
		GroupID: "HandlerReader",

		// MaxWait: 3*time.Second,//wait 3 sec befor recieve
		// StartOffset: kafka.FirstOffset,
	})

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			panic("can't read msg "+err.Error())
		}

		
		

		var m Message

		

		if err := reader.CommitMessages(ctx, msg); err != nil {
            log.Fatalf("could not commit message: %v", err)
        } else {
			json.Unmarshal(msg.Value,&m)
			m.IsChecked = true // handle message
			res,err := json.Marshal(&m)
			
			if err != nil{
				log.Fatal(err.Error())
			} else {

				ProduceMsg(ctx,res,string(msg.Key))
				fmt.Println("recieved: ",string(msg.Value))
				log.Println("commit ",string(msg.Value))
			}

			
		}
	}
}



func main() {
	
	ctx := context.Background()

	consume(ctx)
}