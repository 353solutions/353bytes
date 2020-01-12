package main

import (
	"fmt"
	"log"
)

type Server struct {
	port    int
	verbose bool
}

type Option func(*Server)

func WithVerbose(s *Server) {
	s.verbose = true
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func NewServer(options ...Option) (*Server, error) {
	s := &Server{
		port:    8080,
		verbose: false,
	}

	for _, opt := range options {
		opt(s)
	}

	return s, nil
}

func main() {
	s, err := NewServer(WithVerbose, WithPort(8000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", s)
}
