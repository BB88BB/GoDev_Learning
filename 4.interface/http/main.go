package main

import (
	"fmt"
	"io"
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

	// Method 1
	/*
		//         type   #element
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	// Method 2
	io.Copy(os.Stdout, resp.Body)

}
