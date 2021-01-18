package remote

import (
	"testing"
	"time"
)

func Test(t *testing.T) {

	// server 1
	rpc := newGoRPC()
	rpc.setOnCalled(func(methodName string, msg string) string {
		if methodName == "get1" {
			return "1"
		}
		return "2"
	})
	rpc.start("0.0.0.0", 10000)

	time.Sleep(time.Second)

	// server 2
	rpc1 := newGoRPC()
	msg := rpc1.call("0.0.0.0", 10000, "get1", "")
	if msg != "1" {
		t.Fatal("expected 1 but get ", msg)
	}
}