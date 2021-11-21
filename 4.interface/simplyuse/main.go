package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// englishBot with getGreeting() string
// -> englishBot is an member of Bot
func (eb englishBot) getGreeting() string {
	return "Hi There"
}

// spanishBot with getGreeting() string
// -> spanishBot is an member of Bot
func (sb spanishBot) getGreeting() string {
	return "Hola"
}
