package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Nuts struct {
	NutsConn *nats.Conn
}

func NewNuts() *Nuts {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Printf("failed to get nuts connection:%s", err.Error())
	}
	defer nc.Close()
	return &Nuts{NutsConn: nc}
}
