package potato

import (
	"testing"
)

func TestServer(t *testing.T) {
	app := NewApplication()
	app.Start()
}