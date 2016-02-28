package main

import (
	"./qr"
	"io/ioutil"
	"os"
)

func main() {
	c, err := qr.Encode("123456", qr.H)
	if err != nil {
		panic(err)
	}
	pngdat := c.PNG()
	ioutil.WriteFile("test.png", pngdat, os.ModePerm)
}
