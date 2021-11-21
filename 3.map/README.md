# GoDev_Map

## Key -> Value
- All Keys should have the same type.
- All Values should have the same type.

## Declare map
- **map [key] value**
```Go
    // Method 1
    // var :=  map[ key ] value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

    // Method 2
    var colors map[string]string

    // Method 3
    colors := make(map[string]string)
    colors["white"] = "#ffffff"
```

## Delete keys and value in map
- delete(mapName, key)
```Go
    delete(colors, "white")
```

## Iterating over Maps
```Go
    func printMap(c map[string]string) {
        for color, hex := range c {
            fmt.Println("Hex code for", color, "is", hex)
        }
    }
```

## Map vs. Struct
<img src="https://github.com/BB88BB/GoDev_Map/blob/main/Map%20vs.%20Struct.png" width="712" height="304" />