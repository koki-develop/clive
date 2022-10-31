package cmd

import (
	"net"
)

func randomUnusedPort() (int, error) {
	addr, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	if err = addr.Close(); err != nil {
		return 0, err
	}

	return addr.Addr().(*net.TCPAddr).Port, nil
}
