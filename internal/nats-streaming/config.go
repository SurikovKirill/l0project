package nats_streaming

type Config struct {
	clusterId string
	clientId  string
	host      string
	subject   string
}

func NewConfig() *Config {
	return &Config{
		host:      "nats://localhost:4223",
		clusterId: "test-cluster",
		subject:   "test",
		clientId:  "test-pub",
	}
}
