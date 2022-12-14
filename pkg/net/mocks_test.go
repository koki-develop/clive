package net

import (
	"net"

	"github.com/stretchr/testify/mock"
)

type mockNet struct {
	mock.Mock
}

func (m *mockNet) Listen(network, address string) (net.Listener, error) {
	args := m.Called(network, address)
	return args.Get(0).(net.Listener), args.Error(1)
}

type mockListener struct {
	mock.Mock
}

var _ net.Listener = (*mockListener)(nil)

func (m *mockListener) Accept() (net.Conn, error) {
	args := m.Called()
	return args.Get(0).(net.Conn), args.Error(1)
}

func (m *mockListener) Addr() net.Addr {
	args := m.Called()
	return args.Get(0).(net.Addr)
}

func (m *mockListener) Close() error {
	args := m.Called()
	return args.Error(0)
}
