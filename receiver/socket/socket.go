package socket

import (
	"github.com/anchnet/transfer/g"
	log "github.com/cihub/seelog"
	"net"
)

func StartSocket() {
	if !g.Config().Socket.Enabled {
		return
	}

	addr := g.Config().Socket.Listen
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Errorf("net.ResolveTCPAddr fail: %s", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Errorf("listen %s fail: %s", addr, err)
	} else {
		log.Info("socket listening", addr)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Info("listener.Accept occur error:", err)
			continue
		}

		go socketTelnetHandle(conn)
	}
}
