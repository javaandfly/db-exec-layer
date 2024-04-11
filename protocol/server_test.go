package protocol

import (
	"testing"

	"github.com/panjf2000/gnet"
)

func TestServer(t *testing.T) {
	echo := new(echoServer)
	err := gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true))
	if err != nil {
		t.Fatalf("Failed to serve err is: %v", err)
	}
}
