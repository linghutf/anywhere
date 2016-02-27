package main

import (
	"./qr"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func WritePng(pngname, data string) error {
	c, err := qr.Encode(data, qr.H)
	if err != nil {
		return err
	}
	pngdat := c.PNG()
	return ioutil.WriteFile(pngname, pngdat, os.ModePerm)
}

//generate QR picture by server addr (local network)
func GenQRCodeByAddr(ipstrs *[]string, port int) {
	//find local host ip
	index := 0
	for _, ip := range *ipstrs {
		if strings.Contains(ip, "192.168.") {
			addr := fmt.Sprintf("http://%s:%d", ip, port)
			pngname := strconv.Itoa(index) + ".png"
			err := WritePng(pngname, addr)
			if err != nil {
				log.Fatal(err)
			}
			index++
			fmt.Printf("服务器[%s]二维码%s生成!扫一扫访问.\n", addr, pngname)
		}
	}
}

func VisitIps(port int) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Println("Possible visit :")
	index := 1

	ips := make([]string, len(addrs))

	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				ipstr := ip.IP.String()
				ips[index-1] = ipstr
				fmt.Printf("%-5d http://%s:%d\n", index, ipstr, port)
				index++
			}
		}
	}
	GenQRCodeByAddr(&ips, port)
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
