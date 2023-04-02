package log

import (
	"encoding/json"
	"fmt"

	"github.com/DevellSoftware/go-messenger/pkg/messaging"
)

type Tracer struct {
	messenger *messaging.Messenger
}

func NewTracer(messenger *messaging.Messenger) *Tracer {
	return &Tracer{
		messenger: messenger,
	}
}

func (t *Tracer) Trace(topic string, message string, args ...interface{}) {
	messageString := fmt.Sprintf(message, args...)

	msg := map[string]interface{}{
		"message": messageString,
		"topic":   topic,
	}

	msgBytes, _ := json.Marshal(msg)

	t.messenger.Send(msgBytes)
}
