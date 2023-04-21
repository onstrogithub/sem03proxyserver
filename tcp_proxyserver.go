package main

import (
	"io"
	"log"
	"net"
	"sync"
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

	server, err := net.Dial("tcp", "172.17.0.2:5000")
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		io.Copy(server, client)
		wg.Done()
	}()
	go func() {
		io.Copy(client, server)
		wg.Done()
	}()
	wg.Wait()
}
/*
Dette var koden jeg orginalt sett hadde for handleconnection. Metoden over ble delvis gjennomfort av chatgpt. Jeg fikk laget en forbindelse 
og samlet pakker med koden nedenfor, men etter rundt 1 minutt saa begynnte jeg aa faa feilmeldinger. 
Pakkene ble samlet og pushet til github. Har pakkene fra feilmeldingen paa wireshark dersom dette 
blir interessant aa se paa paa et senere tidspunkt.
func handleConnection(client net.Conn) {
	defer client.Close()

	server, err := net.Dial("tcp", "172.20.0.1:6000")
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()

	go io.Copy(server, client)
	io.Copy(client, server)
}
*/
