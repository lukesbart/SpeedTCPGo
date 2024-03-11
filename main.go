package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var port uint16 = 42000

/* func handleConnection(connection net.Conn, cNum int) {
	deck := game.CreateDeckBox()

	connection.Write([]byte(fmt.Sprintf("You are connection: %d", cNum)))

	cards, _ := json.Marshal(deck.Deck)

	connection.Write([]byte(cards))

	for {
		connection.Write([]byte("Test"))
		time.Sleep(5 * time.Second)
		connection.Write([]byte("Test"))
		connection.Close()
		fmt.Printf("Connection %d closed", cNum)
		return
	}
} */

func acknowledgeConnection(connection net.Conn, cNum int) {
	connection.Write([]byte(fmt.Sprintf("Welcome to SpeedTCP!\nYou are client %d.\n", cNum)))
	return
}

func handleroom(room chan net.Conn) {
	time.Sleep(3 * time.Second)

	for player := range room {
		player.Write([]byte(fmt.Sprintf("The game is starting!\n")))
		// player.Write([]byte(fmt.Sprintf("You are player %d\n", i)))
		player.Write([]byte(fmt.Sprintf("Goodbye!\n")))
		player.Close()
	}

}

func main() {
	fmt.Printf("Server started at port %d\n", port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatal(err)
	}

	cCount := 0

	room := make(chan net.Conn, 2)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		clientAddr := conn.RemoteAddr().String()

		fmt.Printf("New connection at %s\n", clientAddr)

		go acknowledgeConnection(conn, cCount)

		cCount++

		if cCount > 1 {
			break
		}
		room <- conn
	}

	// close(room)
	// go handleroom(room)

	time.Sleep(3 * time.Second)
	for player := range room {
		player.Write([]byte(fmt.Sprintf("The game is starting!\n")))
		// player.Write([]byte(fmt.Sprintf("You are player %d\n", i)))
		player.Write([]byte(fmt.Sprintf("Goodbye!\n")))
		player.Close()
	}

	/* for i, client := range clients {
		client.Write([]byte(fmt.Sprintf("The game is starting!\n")))
		client.Write([]byte(fmt.Sprintf("You are player %d\n", i)))
		client.Write([]byte(fmt.Sprintf("Goodbye!\n")))
		client.Close()
	} */
}
