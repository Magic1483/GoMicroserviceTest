package main

//Handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)



func (s *server) HelpHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w,"URL Commands\n/get - return exists messages from database\n"+
		"/insert?text=<some text> - add new message\n/get?id=<id of message> - get message by id")		
	}
}

func (s *server) DeleteHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		s.db.DeleteMessages()
		fmt.Fprint(w,"Delete all messages")		
	}
}

func (s *server) Handle_insert() httprouter.Handle  {
	type Req struct {
		Text string
	}
	
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var req Req
		var m Message

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		fmt.Println(req)
		
		if req.Text != "" {
			last_id := s.db.InsertMessage(req.Text)
			m = Message{Id: last_id ,Text: req.Text,IsChecked: false}

			answer := fmt.Sprintf("record inserted with id %d",last_id)
			fmt.Fprintf(w,answer)
			text,err := json.Marshal(m)

			if err != nil {
				log.Fatal("can't convert message to JSON",err.Error())
			} else {
				s.ProduceMsg(text,string(last_id)) // send msg -->> msg handler
			}
		} else {
			fmt.Fprintf(w,"text is empty")
		}
		


	}
}

func (s *server) Handle_get() httprouter.Handle  {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		messages,err := s.db.GetMessages()

		if err != nil {
			panic(err)
		}
		
		encoder := json.NewEncoder(w)
		err = encoder.Encode(messages)

		if err != nil {
			panic(err)
		}
		

	}
}

