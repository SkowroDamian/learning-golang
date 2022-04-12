package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	// fmt.Println("Hello, world")

	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
