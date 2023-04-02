package main

import (
	"github.com/DevellSoftware/go-messenger/pkg/messaging"
	"github.com/DevellSoftware/metis-engine/pkg/log"
	"github.com/lovoo/goka"
)

func main() {
	var messenger *messaging.Messenger

	messenger = messaging.NewMessenger("localhost:29092", "metis-engine-events", "metis-engine-events", func(ctx goka.Context, msg interface{}) {
		log.Log("Message: %v", msg)
	})

	messenger.Listen()
}
