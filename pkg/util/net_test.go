package util

import (
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestRandomUnusedTCPPort(t *testing.T) {
	addr := &net.TCPAddr{Port: 999}

	mockNet := &mockNet{}
	mockListener := &mockListener{}

	mockNet.On("Listen", "tcp", ":0").Return(mockListener, nil).Once()
	mockListener.On("Close").Return(nil).Once()
	mockListener.On("Addr").Return(addr, nil).Once()

	netListen = mockNet.Listen

	got, err := RandomUnusedTCPPort()

	assert.NoError(t, err)
	assert.Equal(t, 999, got)
	mockNet.AssertExpectations(t)
	mockListener.AssertExpectations(t)
}
