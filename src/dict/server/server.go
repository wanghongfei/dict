package main

import (
	"log"
	"net/http"
	"dict/server/handler"
	"flag"
	"strconv"
)


func main()  {
	targetPort := flag.Int("p", 9000, "指定端口号")
	flag.Parse()


	host := ":" + strconv.Itoa(*targetPort)
	log.Printf("启动server, %s\n", host)

	http.Handle("/query", handler.NewHandler(new(handler.QueryHandler)))


	err := http.ListenAndServe(host, nil)
	if nil != err {
		log.Fatal(err)
	}
}

