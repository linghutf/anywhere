package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	var port = flag.Int("p", 8001, "server listening port.")
	flag.Parse()
	if *port > 65535 || *port <= 1000 {
		fmt.Println("port illegal!")
		return
	}
	ip_str := "0.0.0.0"
	socket_str := ip_str + ":" + strconv.Itoa(*port)
	go func() {
		fmt.Printf("Listening is on %s...\n", socket_str)
	}()
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(socket_str, nil)

}
