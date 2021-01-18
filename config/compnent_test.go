package config

import (
	"testing"
)

func Test (t *testing.T) {
	c := &Compnent{}
	c.Load("./config_example.json")
	c.SetEnv("dev")
	c.SetLocalPeer("gate", "gate0001")

	pi := c.GetLocalPeerInfo()
	if pi.PeerID != "gate0001" {
		t.Fatal("config error", pi)
	}

	m := c.GetAllPeerInfo()
	if m["chat0000"].PeerID != "chat0000" {
		t.Fatal("config error", m)
	}
}