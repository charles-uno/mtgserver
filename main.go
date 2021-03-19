package main

import (
    "fmt"

    "github.com/charles-uno/mtgserver/lib"
)


// Note: we want to be able to run multiple models for the same opening hand.
// That'll be easier if we send the hand and the library into the game state
// constructor separately. Rather than, say, passing in a list of 60 cards and
// having the constructor shuffle and draw.

func main() {

    deck := lib.LoadDeck()

    hand, deck := deck[:7], deck[7:]

    fmt.Println(hand)


}
