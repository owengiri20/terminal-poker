package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/notnil/joker/hand"
	"github.com/notnil/joker/table"
)

// type Table struct {
// 	options Options
// 	seats   []*Player
// 	dealer  hand.Dealer
// 	deck    *hand.Deck
// 	cards   []hand.Card
// 	active  *Player
// 	status  Status
// 	round   Round
// 	button  int
// 	cost    int
// }

// type Player struct {
// 	ID         string
// 	Seat       int
// 	Chips      int
// 	ChipsInPot int
// 	Acted      bool
// 	Folded     bool
// 	AllIn      bool
// 	Cards      []hand.Card
// }

func main() {
	// starting deck
	src := rand.NewSource(int64(randInt()))
	r := rand.New(src)

	// starter table
	// table := &table.Table{}

	deck := hand.NewDealer(r).Deck()

	players := genPlayers(*deck)

	fmt.Println("==== players ===")
	printPlayers(players)

	// // flop
	// flop := deck.PopMulti(3)

	// h1 := hand.New(deck.PopMulti(5))

	// fmt.Println(h1.Cards())
	// fmt.Println(h1.Ranking() - 1)

	// fmt.Println(randInt())

}

func randInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(52)
}

// type Player struct {
// 	ID          string
// 	Name        string
// 	Card1       hand.Card
// 	Card2       hand.Card
// 	ChipBalance float64
// 	Position    int
// }

// type Table struct {
// 	PotBalance     string
// 	DealerPosition int
// 	hand           hand.Card
// 	deck           hand.Deck
// }

// type Player struct {
// 	ID         string
// 	Seat       int
// 	Chips      int
// 	ChipsInPot int
// 	Acted      bool
// 	Folded     bool
// 	AllIn      bool
// 	Cards      []hand.Card
// }

func genPlayers(deck hand.Deck) []table.Player {
	players := []table.Player{}
	for i := 1; i < 10; i++ {
		players = append(players, newPlayer(deck.PopMulti(2), 5000, i))
	}

	return players

}

func newPlayer(cards []hand.Card, chips int, seat int) table.Player {
	return table.Player{
		ID:         uuid.New().String(),
		Seat:       seat,
		Chips:      chips,
		ChipsInPot: 0,
		Acted:      false,
		Folded:     false,
		AllIn:      false,
		Cards:      cards,
	}
}

func printPlayers(players []table.Player) {
	for i, p := range players {
		fmt.Println("palyer no: ", i)
		fmt.Println(p.ID)
		fmt.Println(p.Cards)
		fmt.Println(hand.New(p.Cards).Ranking() - 1)

		fmt.Println()
	}
}
