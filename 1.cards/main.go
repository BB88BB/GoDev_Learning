package main

//var card string      // can initialize variable outside of the function, but can't assign value to it

func main() {

	// Variable initialize and assign value
	/*
		var card string = "Ace of Spades"      // initialize and assign value, same as the line below
		card := "Ace of Spades"       // initialize and assign value
		card = "Five of Dimonds"      // only assign value (variable have to initialize before)
		card := newCard()
	*/

	// Custom type and slice append
	/*
		cards := deck{"Ace of Diamonds", newCard()}
		cards = append(cards, "Six of Spades") //return a new slices, not modify existed slices
		cards.print()
	*/

	// for loop
	/*
		cards := newDeck()
		cards.print()
	*/

	// Multiple return values, Slice Range Syntax
	/*
		hand, remainingDeck := deal(cards, 5)
		hand.print()
		remainingDeck.print()
	*/

	// Type conversion
	/*
		greeting := "hi here"
		fmt.Println([]byte(greeting))
	*/

	// Save file
	/*
		fmt.Println(cards.toString())
		cards.saveToFile("mycards")
	*/

	// Load file
	/*
		cards := newDeckFromFile("mycards")
		cards.print()
	*/

	// Random number generator, Swap two variable
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string { //return type string data
	return "Six of Diamonds"
}
