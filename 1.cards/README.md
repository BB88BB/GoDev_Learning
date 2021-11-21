# GoDev_CARDS

## Variables
- initialize and assign value, only need initialize once
  ```Go
   var card string = "Ace of Spades"
   card := "Ace of Spades"     
  ``` 
- assign (after initialize)
  ```Go
	card = "Five of Dimonds" 
  ```
  
## Slice: array can grow or shrink
- initialize and assign value  
  ```Go
  cards := []string{"Ace of Diamonds", newcard()}
  ```
- append: not modify existed slices, will return a new slice
  ```Go
  cards = append(cards, "Six of Spades")      
  ```
  
## Iteration
- for loop, will generate index and value in each run
  ```Go
  for index, card := range cards{
  
    fmt.Println(card)
  }
  ```
- If have unused variable, use '_' to represent
  ```Go
  for _, suit := range CardSuit{
    ...
  }
  ```
  
  
## Function
- return one argument
  ```Go
  func newCard() string{
    ...
    return "Five of Diamonds"
  }
- return multi argument
  ```Go
  func newCard() (string, string){
    ...
    return "...", "..."
  }
  left, right = newCard()
  
## Go is not OO -> datatype + receiver 
- define datatype
  ```Go
  type deck []string
  ```
- define receiver: like method, belongs to instance
  ```Go 
  func (d deck) print(){
    ...
  }
  cards := deck{"Ace of Diamonds", newcard()}
  cards.print()
  ```
  
## Type conversion
- convert string to byte slice
  ```Go
  []byte("Hi here!")
  ```

## Join []string to string
- e.g. ["1", "2", "3"] -> "1, 2, 3"
- Need to import "strings" first
  ```Go
  import(
      "strings"
  )
  a := strings.Join( []string(d), ',' )
  ```
  
## Write []string to hard drive
- Need to import "io/ioutil" first
- If fail, return error message
  ```Go
  import(
      "strings"
      “io/ioutil"
  )
  ioutil.WriteFile(filename, []byte(d.toString()), 0666)
  // 0666: anyone can R/W
  ```
  
## Read from the hard drive
- Need to import "io/ioutil" first
- If read file is fail, return error message
  ```Go
  import(
      “io/ioutil"
  )
  bs, err := ioutil.ReadFile(filename)
  ```

## Error message and Exit code
- Need to import "os" first
- If nothing went wrong, err value = 'nil'
- Exit program: use **os.Exit**
- Given status code to exit
  - 0: success; non-0: an error
  ```Go
  import "os"
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
  ```

## Random Number Generator
- Need to import "math.rand" first
- Have to set *Rand by source(e.g. time), then can actually random
- Intn(n) will return number between [0, n)
  ```Go
  import(
      "math.rand"
  )
	
  // create seed source and set *Rand by source
  source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

  // *r.Intn* 
  randnum := r.Intn(len(d) - 1)
  ```

## Swap two variable
  ```Go
  d[i], d[j] = d[j], d[i]
  ```

## Make Test
- Create a new file ending in **_test.go**
  ```
   deck_test.go
  ```
  
- Run all test in packages
  ```Go
  go test
  ```