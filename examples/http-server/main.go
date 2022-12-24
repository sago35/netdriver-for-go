package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func handlerChunked(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Transfer-Encoding", "chunked")
	flusher, ok := w.(http.Flusher)
	if !ok {
		log.Fatal("error")
	}

	io.WriteString(w, "Chunk1\n")
	flusher.Flush()

	io.WriteString(w, "Chunk2Chunk2\n")
	flusher.Flush()
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Chunk1\n")
}

func run() error {
	port := flag.String("port", "8080", "port")
	flag.Parse()

	http.HandleFunc("/chunked", handlerChunked)
	http.HandleFunc("/", handler)

	addr := net.JoinHostPort("", *port)
	fmt.Printf("listen %s\n", addr)
	return http.ListenAndServe(addr, nil)
}
