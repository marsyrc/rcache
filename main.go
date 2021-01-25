package main

import (
	"flag"
	"log"
	"rcache/cache"
	"rcache/http"
	"rcache/tcp"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listen()
	http.New(c).Listen()
}