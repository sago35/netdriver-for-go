# netdriver-for-go

This package was created to communicate from Go using `tinygo.org/x/drivers/net`.  

## how to use

```go
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
	var buf [1024 * 10]byte
	net.UseDriver(&netdriver.Driver{
		MaxPacketSize: 4096,
		Debug:         false,
	})
	http.SetBuf(buf[:])

	res, err := http.Get("http://tinygo.org")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%q\n", string(b))
	return nil
}
```

## LICENSE

MIT

## Author

sago35 - <sago35@gmail.com>
