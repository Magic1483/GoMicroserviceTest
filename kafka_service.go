package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// https://www.sohamkamani.com/golang/working-with-kafka/
// kafka UI https://github.com/provectus/kafka-ui

const (
	SendTopic = "server_to_handler"
	RecieveTopic = "handler_to_server"
	Address = "kafka:9092"

)

func CreateTopic(ctx context.Context) {
	_, err := kafka.DialLeader(ctx,"tcp",Address,SendTopic,0)
	if err != nil {
		panic(err.Error())
	}
	_, err = kafka.DialLeader(ctx,"tcp",Address,RecieveTopic,0)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("topics created")
}

func (s *server) ProduceMsg(msg []byte,key string) {
 
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{Address},
		Topic: SendTopic,
	})

	err := w.WriteMessages(*s.kafka_ctx,kafka.Message{
		Key: []byte(key),
		Value: []byte(msg),
	})
	if err != nil {
		panic("can't write msg "+err.Error())
	}
}

func (s *server) consume(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{Address},
		Topic: RecieveTopic,
		MinBytes: 5,
		MaxBytes: 1e6,
		GroupID: "ServerReader",

		// MaxWait: 3*time.Second,//wait 3 sec befor recieve
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
			fmt.Println("recieved: ",string(msg.Value))
			log.Println("commit ",string(msg.Value))
			
			json.Unmarshal(msg.Value,&m) // create Struct from kafka msg
			err = s.db.UpdateMessage(&m)
			if err != nil {
				log.Fatal("message not updated",err.Error())
			}
		}


	}
}


func (s *server) SetupKafka()  {
	ctx := context.Background()

	CreateTopic(ctx)
	go s.consume(ctx)

	s.kafka_ctx = &ctx
}