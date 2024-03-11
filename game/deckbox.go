package game

var suites []string = []string{"Spades", "Diamonds", "Clubs", "Hearts"}
var cardNames []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
var values []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

type Card struct {
	Suite    string
	CardName string
	Value    int
}

type DeckBox struct {
	Deck []Card
}

func CreateDeckBox() DeckBox {
	newDeckBox := DeckBox{}

	for _, suite := range suites {
		for i, cardName := range cardNames {
			newDeckBox.Deck = append(newDeckBox.Deck, Card{suite, cardName, values[i]})
		}
	}

	return newDeckBox
}
