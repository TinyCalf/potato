package main

import (
	"potato/piface"
	"potato/common"
	"potato/remote"
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

	remote := app.GetComponent("Remote").(remote.ICompnent)
	remote.SetAddress("0.0.0.0", 10000)
	remote.AddPeer(0, "0.0.0.0", 10001)
	remote.RegistMethod("echo", func(msg string) string {
		return msg
	})

	app.Start()

	// for {
	// 	resp := remote.Call(1, "echo", "hello world!")
	// 	fmt.Printf("Remote Response:%s\n", resp)
	// 	time.Sleep(time.Second)
	// }

}