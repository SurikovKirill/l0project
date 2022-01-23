package main

import (
	"github.com/BurntSushi/toml"
	"l0project/internal/apiserver"
	"l0project/internal/cache"
	nats_streaming "l0project/internal/nats-streaming"
	"log"
	"time"
)

func main() {
	// Initialize cache service
	cst := cache.New(5*time.Minute, 10*time.Minute)
	// Initialize web server
	c := apiserver.NewConfig()
	_, err := toml.DecodeFile("configs/apiserver.toml", c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c)
	s := apiserver.New(c, cst)

	// Initialize Nats subscriber
	nsc := nats_streaming.NewConfig()
	_, err = toml.DecodeFile("configs/apiserver.toml", nsc)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nsc)

	ns := nats_streaming.New(nsc, cst)
	log.Println("Starting nats")
	go ns.Start()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
