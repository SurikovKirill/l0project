package nats_streaming

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

type Stan struct {
	client *stan.Conn
	config *Config
}

func New(config *Config) *Stan {
	return &Stan{
		config: config,
	}
}

func (st *Stan) Start() error {
	// Connect to nats-streaming
	sc, err := stan.Connect(st.config.clusterId, st.config.clientId, stan.NatsURL(st.config.host),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		return err
	}

	// Subscribe
	sub, err := sc.Subscribe(st.config.subject, func(msg *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(msg.Data))
	}, stan.DeliverAllAvailable())
	if err != nil {
		return err
	}

	// Unsubscribe
	if err := sub.Close(); err != nil {
		return err
	}

	// Close connection
	if err := sc.Close(); err != nil {
		return err
	}
	return nil
}
