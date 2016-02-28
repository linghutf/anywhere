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
	"os/exec"
	"path/filepath"
	"runtime"
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

/*
deprecated
show img by local default webbrowser
*/ /*
func ShowPng(pngname string) {
	filepath, err := filepath.Abs(pngname)
	if err != nil {
		log.Fatal(err)
		return
	}
	fileaddr := fmt.Sprintf("file:///%s", filepath)
	webbrowser.Open(fileaddr)
}
*/
/*
 * show image by Qt
 * addfinfo : server addr
 * filename : QR png filename
 * platform : environment
 */
func ShowImage(addrinfo, filename string) {
	var binfilename string
	//judge platform
	switch runtime.GOOS {
	case "windows":
		binfilename, _ = filepath.Abs("./show/bin/win/SHowQRImg.exe")
	default:
		binfilename, _ = filepath.Abs("./show/bin/SHowQRImg")
	}
	file, _ := filepath.Abs(filename)
	// fmt.Printf("%s %s %s\n", binfilename, addrinfo, file)
	// run commad
	c := exec.Command(binfilename, addrinfo, file)
	c.Run()
}

//generate QR picture by server addr (local network)
//return map[addrinfo]pngfilename
func GenQRCodeByAddr(ipstrs *[]string, port int) *map[string]string {
	//find local host ip
	infos := make(map[string]string, len(*ipstrs))
	index := 0
	for _, ip := range *ipstrs {
		if strings.Contains(ip, "192.168.") {
			addr := fmt.Sprintf("http://%s:%d", ip, port)
			//pngname := strconv.Itoa(index) + ".png"
			pngname := fmt.Sprintf("%d-addr.png", index)
			err := WritePng(pngname, addr)
			if err != nil {
				log.Fatal(err)
			}
			infos[addr] = pngname
			index++
			fmt.Printf("服务器[%s]二维码生成!扫一扫当前目录下[%s]访问.\n", ip, pngname)
			// ShowPng(pngname)
		}
	}
	return &infos
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
	infos := GenQRCodeByAddr(&ips, port)
	//show images
	for addr, filename := range *infos {
		go ShowImage(addr, filename)
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
