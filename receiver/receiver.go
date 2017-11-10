package receiver

import (
	"github.com/anchnet/transfer/receiver/rpc"
	"github.com/anchnet/transfer/receiver/socket"
)

func Start() {
	go rpc.StartRpc()
	go socket.StartSocket()
}
