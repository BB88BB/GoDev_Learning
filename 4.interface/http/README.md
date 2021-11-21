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
<img src="https://github.com/BB88BB/GoDev_Learning/blob/main/4.interface/http/response__struct.png" width="944" height="364" />

## Reader interface
- reason  
<img src="https://github.com/BB88BB/GoDev_Learning/blob/main/4.interface/http/reason.png" width="836" height="431" />

- Reader interface  
<img src="https://github.com/BB88BB/GoDev_Learning/blob/main/4.interface/http/reader_interface.png" width="825" height="427" />

- implemet Reader
<img src="https://github.com/BB88BB/GoDev_Learning/blob/main/4.interface/http/read_implementation.png" width="694" height="370" />

## get body field
```Go
bs := make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
```