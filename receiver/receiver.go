package receiver

import (
	"github.com/51idc/transfer/receiver/rpc"
	"github.com/51idc/transfer/receiver/socket"
)

func Start() {
	go rpc.StartRpc()
	go socket.StartSocket()
}
