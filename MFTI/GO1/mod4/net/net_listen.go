package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	name := conn.RemoteAddr().String() // set name as ip

	// show other participants in server name of connected user
	fmt.Printf("%+v connected\n", name)
	// conn.Write() write mssg only for current user
	conn.Write([]byte("Hello, " + name + "\n\r"))

	defer conn.Close()

	// scanner for input reading
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Exit" { // stop listening for this user
			conn.Write([]byte("Bye\n\r"))
			fmt.Println(name, "disconnected")
			break
		} else if text != "" {
			fmt.Println(name, "enters", text)
			conn.Write([]byte("You enter " + text + "\n\r"))
		}
	}
}

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// listen server input
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		// handleConnection must be run in async to work with few connections
		go handleConnection(conn)
	}
}
