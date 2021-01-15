package main

import (
	"potato/piface"
	"potato/common"
	"potato"
	"fmt"
)

// EchoRouter ..
type EchoRouter struct{
	common.BaseRouter
}

// Handle ..
func (r *EchoRouter) Handle(session piface.ISession, msg piface.IMessage){
	err := session.Send(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	app := potato.NewApplication()

	handler := app.GetComponent("HandlerService").(piface.IHandlerService)
	handler.AddRouter(1, &EchoRouter{})

	app.Start()
}