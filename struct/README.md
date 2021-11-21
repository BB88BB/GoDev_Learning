# GoDev_struct

## Define struct type
```Go
type person struct {
	firstname string
	lastname  string
}
```

## Update struct value
- Without assign value -> assign zero value ("")
```Go
alex := person{firstname: "Alex", lastname: "Anderson"}
	
var vivian, tom  person
tom.firstname = "tom"
tom.lastname = "Anderson"
```
- %v, %+v, %#v different way to print
```Go
fmt.Println(alex)
// {Alex Anderson}

fmt.Printf("%+v", vivian)
// {firstname: lastname:}%

fmt.Printf("%#v", tom)
// main.person{firstname:"Tom", lastname:"Anderson"}%
```

## Embedded struct
- Define embedded struct
```Go
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}
```
- Assign value to embedded struct
```Go
jim := person{
	firstname: "Jim",
	lastname:  "Andrew",
	contact: contactInfo{
		email:   "jim@gmail.com",
		zipCode: 94000,
	},    // Very important
}
```
- Actually, can only define like this
```Go
type person struct {
    ...
	contactInfo
}

jim := person{
	...
	contactInfo: contactInfo{
		...
	},
}
```

# Pointer in Go
- &variable : get the access to the **memory address**
    `Turn value into address with &value`
- *pointer  : get the **value** of this memory address point at
    `Turn address into value with *address`

```Go
jimPointer := &jim
jimPointer.updateName("jimmy")
jim.print()

// ptrToPerson *person -> only pointer of person type can call this receiver
func (ptrToPerson *person) updateName(newfirstname string){

    // get the value of this pointer
	(*ptrToPerson).firstname = newfirstname
}
```

- Pointer shortcut, will allow to call receiver *person
```Go
jim.updateName("jimmy")
jim.print()
```
## Reference vs. Value types
- Reference type: Don't worry about pointers
    - `slice`, `maps`, `channels`, `pointers`, `functions`
    - When create a slice, Go will automactically create: 
        1. `an array`
        2. `structure`: length, capacity, reference to the underlying array 
- Value type: Use pointer to change these things in a function
    - `int`, `float`, `string`, `bool`, `structs`

## Go always pass by value
![image](https://github.com/BB88BB/GoDev_struct/blob/master/always%20pass%20by%20value.jpeg)