package nats_streaming

import "l0project/internal/store"

type Config struct {
	ClusterId string
	ClientId  string
	Host      string
	Subject   string
	Store     *store.Config
}

func NewConfig() *Config {
	return &Config{
		Host:      "nats://localhost:4223",
		ClusterId: "test-cluster",
		Subject:   "test",
		ClientId:  "test-sub",
		Store:     store.NewConfig(),
	}
}
