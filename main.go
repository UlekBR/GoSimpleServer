package main

import (
	"github.com/UlekBR/GoSimpleServer/server"
)

func main() {
	handler := &server.Handler{}

	handler.GetAndPost("/", func(rw *server.ResponseManager) {
		rw.WriteString("hello world " + rw.RequestURI)
	})

	handler.Get("/get", func(rw *server.ResponseManager) {
		rw.WriteString("hello world " + rw.Host)
	})

	handler.Post("/post", func(rw *server.ResponseManager) {
		rw.WriteString("hello world")
	})

	server.Start("0.0.0.0", 8081, "Servidor rodando em http://{host}:{port}\n")

}
