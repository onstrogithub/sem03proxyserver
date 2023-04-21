package main

import (
	"io"
	"log"
	"net"
)

func main() {
	proxyServer, err := net.Listen("tcp", "0.0.0.0:6000")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s", proxyServer.Addr().String())

	for {
		conn, err := proxyServer.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(client net.Conn) {
	defer client.Close()

	server, err := net.Dial("tcp","172.20.0.1:6000") //"172.17.0.2:5000")
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()

	go io.Copy(server, client)
	io.Copy(client, server)
}

