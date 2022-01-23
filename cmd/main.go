package main

import (
	"github.com/BurntSushi/toml"
	"l0project/internal/apiserver"
	"l0project/internal/cache"
	natsstreaming "l0project/internal/nats-streaming"
	"log"
	"time"
)

func main() {
	// Initialize cache service
	cch := cache.New(5*time.Minute, 10*time.Minute)
	// Initialize web server
	c := apiserver.NewConfig()
	_, err := toml.DecodeFile("configs/apiserver.toml", c)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(c, cch)
	// Initialize Nats subscriber
	nsc := natsstreaming.NewConfig()
	_, err = toml.DecodeFile("configs/apiserver.toml", nsc)
	if err != nil {
		log.Fatal(err)
	}
	ns := natsstreaming.New(nsc, cch)
	log.Println("Starting nats")
	go func() {
		if err := ns.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Starting web server")
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
