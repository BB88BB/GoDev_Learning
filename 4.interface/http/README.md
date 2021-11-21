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
<img src="https://github.com/BB88BB/GoDev_Learning/blob/main/4.interface/http/response__struct.png" width="472" height="182" />

## Reader interface
- Reader interface: read information from outside data source into application
- reason  
<img src="https://github.com/BB88BB/GoDev_Learning/blob/main/4.interface/http/reason.png" width="218" height="215" />

- Reader interface  
<img src="https://github.com/BB88BB/GoDev_Learning/blob/main/4.interface/http/reader_interface.png" width="412" height="213" />

- implemet Reader
<img src="https://github.com/BB88BB/GoDev_Learning/blob/main/4.interface/http/read_implementation.png" width="347" height="185" />

## Get body field - use Reader
```Go
//         type   #element
bs := make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
```

## Get body field - use Writer
- Writer interface: take info and sent out
- io.Copy(dst Writer, source Reader)
```Go
io.Copy(os.Stdout, resp.Body)
```
## Custom Writer
```Go
type logWriter struct{}

lw := logWriter{}
io.Copy(lw, resp.Body)

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("len of bs", len(bs))
	return len(bs), nil
}
```