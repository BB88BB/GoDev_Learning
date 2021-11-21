# GoDev_Interface

## Declare Interface
- New custom type "bot"
- If you are a type in program, and with a function **getGretting() string**  
You are now an member of type "bot"
- Each member in "bot" can call function **printGreeting**

```Go
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

// englishBot associate with getGreeting() string 
// -> englishBot is an member of Bot
func (eb englishBot) getGreeting() string {
	return "Hi There"
}

// spanishBot associate with getGreeting() string 
// -> spanishBot is an member of Bot
func (sb spanishBot) getGreeting() string {
	return "Hola"
}
```

## Concrete Type vs. Interface Type
- Concrete Type: `map`, `struct`, `int`, `string`, `englishBot`...
- Interface Type:  `bot`

## Extra note
- Interface are not generic type
- Interface are 'implicit'  
- Interface are a contract to help us manage types  
garbage in -> garbage out

