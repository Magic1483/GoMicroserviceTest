package main

import "github.com/julienschmidt/httprouter"

func (s *server) setupRoutes()  {
	s.router = httprouter.New()

	s.router.GET("/get",s.Handle_get())
	s.router.GET("/",s.HelpHandler())
	s.router.POST("/insert",s.Handle_insert())
}