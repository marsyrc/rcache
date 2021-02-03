package tcp

import (
	"net"
	"rcache/server/cache"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	l, e := net.Listen("tcp", ":12345")
	if e != nil {
		panic(e)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			panic(e)
		}
		go s.process(c)
	}
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
