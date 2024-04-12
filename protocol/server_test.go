package protocol

import (
	"testing"

	"github.com/panjf2000/gnet"
)

func TestServerPing(t *testing.T) {
	echo := new(echoServer)
	err := gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true))


	if err != nil {
		t.Fatalf("err : %v", err)
	}
}
