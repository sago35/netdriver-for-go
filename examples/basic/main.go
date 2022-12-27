package main

import (
	"fmt"
	"io"
	"log"

	"github.com/sago35/netdriver-for-go"
	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/net/http"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var buf [1024 * 100]byte
	net.UseDriver(&netdriver.Driver{
		MaxPacketSize: 0,
		Debug:         false,
	})
	http.SetBuf(buf[:])

	res, err := http.Get("http://httpbin.org/ip")
	//res, err := http.Get("https://httpbin.org/ip")
	//res, err := http.Get("http://tinygo.org")
	//res, err := http.Get("https://tinygo.org")
	//res, err := http.Get("http://localhost:8080")
	//res, err := http.Get("http://localhost:8080/chunked")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(b))
	return nil
}
