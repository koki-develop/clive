package util

import "net"

func RandomUnusedTCPPort() (int, error) {
	addr, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	if err = addr.Close(); err != nil {
		return 0, err
	}

	return addr.Addr().(*net.TCPAddr).Port, nil
}
