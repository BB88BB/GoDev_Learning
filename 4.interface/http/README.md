# http

## Get(URL)
- Neet to import "net/http"
- resp: include status, protocol, body...
```Go
import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Errpr: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
```

## http Response struct


## Reader interface
- reason  

- Reader interface  

- implemet Reader

## get body field
```Go
bs := make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
```