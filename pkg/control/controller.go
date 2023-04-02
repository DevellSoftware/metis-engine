package control

import "github.com/DevellSoftware/go-messenger/pkg/messaging"

type Controller struct {
	messenger *messaging.Messenger
}

func NewController(messenger *messaging.Messenger) *Controller {
	return &Controller{
		messenger: messenger,
	}
}
