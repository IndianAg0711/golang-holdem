package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
	"golang.org/x/exp/slices"
)

type Card struct {
	Code  string
	Value string
	Suit  string
}

type Deck struct {
	Success   bool
	Deck_id   string
	Remaining int
	Shuffled  bool
	Cards     []Card
}

var PlayersHand []Card

var DealersHand []Card

var currentDeck Deck

var cardShapes = map[string]string{
	"CLUBS":    "♣",
	"DIAMONDS": "♦",
	"HEARTS":   "♥",
	"SPADES":   "♠",
}

func main() {
	welcomeMessage()

	fmt.Println("Requesting new deck...")

	getNewDeck()
	getPlayersHand()
	getDealersHand()
	// getHand()
	displayHand()
}

func welcomeMessage() {
	color.Cyan("================ Welcome to Texas Hold'em! ================")
	fmt.Println("Press any key to deal a new hand...")

	var any string
	fmt.Scanln(&any)
}

func getNewDeck() {
	resp, err := http.Get("https://deckofcardsapi.com/api/deck/new/shuffle/?deck_count=1")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, &currentDeck)
}

func getPlayersHand() {
	card_url := fmt.Sprintf("https://deckofcardsapi.com/api/deck/%s/draw/?count=2", currentDeck.Deck_id)
	resp, err := http.Get(card_url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, &currentDeck)
	PlayersHand = currentDeck.Cards
	currentDeck.Cards = nil
}

func getDealersHand() {
	card_url := fmt.Sprintf("https://deckofcardsapi.com/api/deck/%s/draw/?count=5", currentDeck.Deck_id)
	resp, err := http.Get(card_url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, &currentDeck)
	DealersHand = currentDeck.Cards
}

func getHand() {
	card_url := fmt.Sprintf("https://deckofcardsapi.com/api/deck/%s/draw/?count=5", currentDeck.Deck_id)
	resp, err := http.Get(card_url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, &currentDeck)
}

func displayHand() {
	color.Green("================ Player's Hand ================")
	for i := 0; i < len(PlayersHand); i++ {
		card := PlayersHand[i]
		redCards := []string{"DIAMONDS", "HEARTS"}
		if slices.Contains(redCards, card.Suit) {
			color.Red("%s %s", card.Value, cardShapes[card.Suit])
		} else {
			fmt.Printf("%s %s\n", card.Value, cardShapes[card.Suit])
		}
	}

	color.Yellow("Press any key to show dealers hand...")

	var any string
	fmt.Scanln(&any)

	color.Blue("================ Dealer's Hand ================")
	for i := 0; i < len(DealersHand); i++ {
		card := DealersHand[i]
		redCards := []string{"DIAMONDS", "HEARTS"}
		if slices.Contains(redCards, card.Suit) {
			color.Red("%s %s", card.Value, cardShapes[card.Suit])
		} else {
			fmt.Printf("%s %s\n", card.Value, cardShapes[card.Suit])
		}
	}
}
