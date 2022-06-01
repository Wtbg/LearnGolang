package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: get to url %v\n", err)
			os.Exit(1)
		}
		// bodycontent, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: read body %v\n", err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", bodycontent
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}
