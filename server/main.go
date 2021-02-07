package server

import (
	"flag"
	"log"
	"rcache/server/cache"
	"rcache/server/http"
	"rcache/server/tcp"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
