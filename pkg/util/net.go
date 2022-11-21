package util

import "net"

var netListen = net.Listen

func RandomUnusedTCPPort() (int, error) {
	addr, err := netListen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	if err = addr.Close(); err != nil {
		return 0, err
	}

	return addr.Addr().(*net.TCPAddr).Port, nil
}

type INetListener interface {
	RandomUnusedTCPPort() (int, error)
}

type NetListener struct {
	listen func(network, address string) (net.Listener, error)
}

var _ INetListener = (*NetListener)(nil)

func NewNetListener() *NetListener {
	return &NetListener{
		listen: net.Listen,
	}
}

func (l *NetListener) RandomUnusedTCPPort() (int, error) {
	addr, err := l.listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	if err = addr.Close(); err != nil {
		return 0, err
	}

	return addr.Addr().(*net.TCPAddr).Port, nil
}
