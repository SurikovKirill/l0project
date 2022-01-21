package main

import nats_streaming "l0project/internal/nats-streaming"

func main() {
	c := nats_streaming.NewConfig()
	s := nats_streaming.New(c)
	s.Start()

}
