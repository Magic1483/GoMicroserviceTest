package main

//Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)



func (s *server) HelpHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w,"URL Commands\n/get - return exists messages from database\n"+
		"/insert?text=<some text> - add new message\n/get?id=<id of message> - get message by id")		
	}
}

func (s *server) Handle_insert() httprouter.Handle  {
	type Req struct {
		Text string
	}
	
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var req Req
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		fmt.Println(req)
		
		if req.Text != "" {
			err = s.db.InsertMessage(req.Text)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w,"record inserted")
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

