package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type room struct {
	players []player
	status  chan struct{}
	game    *game
}

type player struct {
	conn  net.Conn
	num   int
	name  string
	cards []card
}

func main() {
	server, err := net.Listen("tcp", ":8080")
	defer server.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting server at :8080")

	status := make(chan struct{})
	newRoom := room{status: status}

	for len(newRoom.players) < 2 {
		newConn, err := server.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		newConn.Write([]byte("Welcome to SpeedTCP!\n"))
		newPlayer := player{conn: newConn, num: len(newRoom.players)}

		newRoom.players = append(newRoom.players, newPlayer)
	}

	go newRoom.handleRoom()

	<-status
	fmt.Println("Game finished")
}

func (r *room) handleRoom() {
	r.broadcast([]byte("Starting game\n"))
	game := newGame()
	r.game = &game

	wg := sync.WaitGroup{}

	for _, player := range r.players {
		wg.Add(1)

		player.cards = game.cards[0:5]
		game.cards = game.cards[5:]

		go r.handlePlayer(player, &wg)
	}

	wg.Wait()

	r.status <- struct{}{}
}

func (r *room) handlePlayer(player player, wg *sync.WaitGroup) {
	player.conn.Write([]byte(fmt.Sprintf("You are player %d\n", player.num)))
	player.conn.Write([]byte("Enter your name: "))
	buf := make([]byte, 1024)
	n, err := player.conn.Read(buf)
	if err != nil || n == 0 {
		player.name = fmt.Sprintf("%d", player.conn.RemoteAddr())
	}
	playerNameRead := fmt.Sprintf("%s", buf[:n])
	player.name = playerNameRead[:n-2]

	if player.num == 0 {
		r.game.player1Hand = player.cards

	} else {

		r.game.player2Hand = player.cards
	}

	gs := r.game.createGameMessage(player, *r)

	player.conn.Write([]byte(fmt.Sprintf("%v\n", gs)))

	buf = make([]byte, 1024)
	n, err = player.conn.Read(buf)

	if err != nil {
		fmt.Printf("Error writing to client %v", err)
	}

	msg := buf[:n]

	r.broadcast(msg)

	wg.Done()
}

func (r *room) broadcast(msg []byte) {
	for _, player := range r.players {
		player.conn.Write(msg)
	}
}
