package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
)

var suites []string = []string{"Spades", "Diamonds", "Clubs", "Hearts"}
var cardNames []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
var values []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

type card struct {
	Name  string `json:"name"`
	Suite string `json:"suite"`
	Value int    `json:"value"`
}

type gameState string

const (
	SETUP    gameState = "SETUP"
	READY              = "READY"
	PROGRESS           = "PROGRESS"
	FINISHED           = "FINISHED"
)

type game struct {
	// Initial container for cards, emptied when game setup
	cards []card

	draw1 []card
	draw2 []card

	discard1 []card
	discard2 []card

	player1Hand []card
	player2Hand []card

	gameStatus gameState
}

type gameMessage struct {
	State gameState `json:"gameState"`

	PlayerName string `json:"playerName"`
	PlayerNum  int    `json:"playernum"`

	OpponentName string `json:"opponentName"`

	PlayerHand []card `json:"playerHand"`

	Draw1Pile []card `json:"draw1Pile"`
	Draw2Pile []card `json:"draw2Pile"`

	Discard1Pile []card `json:"discard1Pile"`
	Discard2Pile []card `json:"discard2Pile"`
}

func (g *game) createGameMessage(player player, room room) string {
	ngm := gameMessage{
		State:        g.gameStatus,
		PlayerName:   player.name,
		PlayerNum:    player.num,
		PlayerHand:   player.cards,
		Draw1Pile:    g.draw1,
		Draw2Pile:    g.draw2,
		Discard1Pile: g.discard1,
		Discard2Pile: g.discard2,
	}

	if player.num == 0 {
		ngm.OpponentName = room.players[1].name
	} else {
		ngm.OpponentName = room.players[0].name
	}

	ngmJson, _ := json.Marshal(ngm)

	ngmText := string(ngmJson)

	fmt.Printf("%v", string(ngmText))

	return ngmText
}

func newGame() game {
	newDeck := createDeck()
	draw1 := newDeck[0:19]
	draw2 := newDeck[20:40]

	newDeck = newDeck[41:]

	return game{cards: newDeck, draw1: draw1, draw2: draw2, gameStatus: SETUP}
}

func createDeck() []card {
	deck := []card{}

	for _, suite := range suites {
		for i, cardName := range cardNames {
			deck = append(deck, card{suite, cardName, values[i]})
		}
	}

	// Fisher-Yates shuffle
	for i := range len(deck) - 2 {
		j := rand.IntN(len(deck))
		tmp := deck[i]
		deck[i] = deck[j]
		deck[j] = tmp
	}

	return deck
}
