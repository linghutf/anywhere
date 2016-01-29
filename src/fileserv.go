package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func VisitIps(port int) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Println("Possible visit :")
	index := 1
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				fmt.Printf("%-5d http://%s:%d\n", index, ip.IP.String(), port)
				index++
			}
		}
	}
}

func main() {
	port := *flag.Int("p", 8001, "server listening port.")
	flag.Parse()
	if port > 65535 || port <= 1000 {
		fmt.Println("port illegal!")
		return
	}

	socket_str := net.JoinHostPort("0.0.0.0", strconv.Itoa(port))
	go func() {
		log.Printf("Listening is on %s...\n", socket_str)
		VisitIps(port)
	}()
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(socket_str, nil)
}
