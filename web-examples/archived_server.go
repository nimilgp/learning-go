package archivemain

import (
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
	"time"
)

type httpRequestParts struct {
	method string
	target string
	hostname string
	useragent string
}

func rootHandle(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	os.Exit(1)
}

func echoHandle(conn net.Conn, msg string) {
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Content-Type: text/plain\r\n"))
	conn.Write([]byte("Content-Length: " + strconv.Itoa(len(msg)) +"\r\n\r\n"))
	conn.Write([]byte(msg))
}

func userAgentHandle(conn net.Conn, useragent string) {
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Content-Type: text/plain\r\n"))
	conn.Write([]byte("Content-Length: " + strconv.Itoa(len(useragent)) +"\r\n\r\n"))
	conn.Write([]byte(useragent))
}

func defaultHandle(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	os.Exit(1)
}

func conHandle(conn net.Conn) { 
	fmt.Println("processing")
	time.Sleep(1*time.Second)
	defer conn.Close()
	buf := make([]byte, 1024)
	var hrp httpRequestParts
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error")
		return
	}
	splitLines := strings.Split(string(buf), "\r\n")
	splitLine1 := strings.Split(splitLines[0], " ")
	splitLine2 := strings.Split(splitLines[1], " ")
	splitLine3 := strings.Split(splitLines[2], " ")
	hrp.method = splitLine1[0]
	hrp.target = splitLine1[1]
	hrp.hostname = splitLine2[1]
	hrp.useragent = splitLine3[1]
	if hrp.target == "/" {
		rootHandle(conn)
	} else if strings.HasPrefix(hrp.target, "/echo/") {
		echoHandle(conn, hrp.target[6:])
	} else if strings.HasPrefix(hrp.target, "/user-agent") {
		userAgentHandle(conn, hrp.useragent)
	} else {
		defaultHandle(conn)
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()
	i:=1
	for {
		conn, err := l.Accept()
		fmt.Println(i)
		i += 1
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			conn.Close()
			break
		}
		go conHandle(conn)
	}
}
