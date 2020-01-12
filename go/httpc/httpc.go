package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
)

var reqTemplate = `GET / HTTP/1.1
Host: %s
Connection: Close

`

func main() {
	host, port := "golang.org", 443
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	request := fmt.Sprintf(reqTemplate, host)
	_, err = conn.Write([]byte(request))
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got %d bytes\n", buf.Len())
	size := 600
	if size >= buf.Len() {
		size = buf.Len() - 1
	}
	fmt.Printf("%s\n", buf.String()[:size])
}
