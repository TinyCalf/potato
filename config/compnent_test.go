package config

import (
	"testing"
)

func Test (t *testing.T) {
	c := &Compnent{}
	c.Load("./config_example.json")
	t.Fatal(c.Config)
}