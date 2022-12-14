package net

import "net"

type IListener interface {
	RandomUnusedTCPPort() (int, error)
}

type Listener struct {
	listen func(network, address string) (net.Listener, error)
}

var _ IListener = (*Listener)(nil)

func NewListener() *Listener {
	return &Listener{
		listen: net.Listen,
	}
}

func (l *Listener) RandomUnusedTCPPort() (int, error) {
	addr, err := l.listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	if err = addr.Close(); err != nil {
		return 0, err
	}

	return addr.Addr().(*net.TCPAddr).Port, nil
}
