package main

import (
	"flag"
	"log"
)

var (
	ip = flag.String("ip", ":8000", "The IP-Address on which the server listens")
)

func main() {
	flag.Parse()
	s := NewServer()
	if err := s.Start(*ip); err != nil {
		log.Fatalln(err)
	}
}

