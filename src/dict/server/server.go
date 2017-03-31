package main

import (
	"log"
	"net/http"
	"dict/server/handler"
)

func main()  {
	log.Println("启动server")

	http.Handle("/query", handler.NewHandler(handler.QueryHandler{}))

	err := http.ListenAndServe(":9000", nil)
	if nil != err {
		log.Fatal(err)
	}
}

