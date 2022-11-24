package main

// Kommentare folgen

import (
	"fmt"
	"math/rand"
	"time"
)

type PlayCard struct {
	Colour string
	Value  string
}

func (card PlayCard) String() string {
	return fmt.Sprintf("%s of %s", card.Value, card.Colour)
}

type Deck []PlayCard

func (deck Deck) String() string {
	result := ""
	for _, card := range deck {
		result += card.String() + "\n"
	}
	return result
}

func (deck Deck) Shuffle() {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func (deck *Deck) getLast() PlayCard {
	return (*deck)[len(*deck)-1]
}

func (deck *Deck) removeLast() {
	*deck = (*deck)[:len(*deck)-1]
}

func (deck *Deck) Draw() PlayCard {
	result := deck.getLast()
	deck.removeLast()
	return result
}

func NewDeck() Deck {
	colours := []string{"clubs", "spades", "hearts", "diamonds"}
	values := []string{"7", "8", "9", "10", "jack", "queen", "king", "ace"}

	deck := Deck{}
	for _, c := range colours {
		for _, v := range values {
			deck = append(deck, PlayCard{c, v})
		}
	}
	return deck
}

func main() {
	rand.Seed(time.Now().UnixNano())

	d := NewDeck()

	d.Shuffle()

	hand1 := Deck{}
	hand1 = append(hand1, d.Draw())

	fmt.Println(d)
	fmt.Println(hand1)
}
