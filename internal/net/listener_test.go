package net

import (
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestNewListener(t *testing.T) {
	l := NewListener()

	expected, actual := reflect.ValueOf(net.Listen).Pointer(), reflect.ValueOf(l.listen).Pointer()

	assert.NotNil(t, l)
	assert.Equal(t, expected, actual)
}

func TestListener_RandomUnusedTCPPort(t *testing.T) {
	type mocks struct {
		net      *mockNet
		listener *mockListener
	}
	tests := []struct {
		mocks   *mocks
		want    int
		wantErr bool
	}{
		{
			func() *mocks {
				mockListener := &mockListener{}
				mockListener.On("Close").Return(nil)
				mockListener.On("Addr").Return(&net.TCPAddr{Port: 999}, nil)
				mockNet := &mockNet{}
				mockNet.On("Listen", "tcp", ":0").Return(mockListener, nil)
				return &mocks{mockNet, mockListener}
			}(),
			999,
			false,
		},
		{
			func() *mocks {
				mockListener := &mockListener{}
				mockNet := &mockNet{}
				mockNet.On("Listen", "tcp", ":0").Return((*net.TCPListener)(nil), errors.New("SOMETHING_WRONG"))
				return &mocks{mockNet, mockListener}
			}(),
			0,
			true,
		},
		{
			func() *mocks {
				mockListener := &mockListener{}
				mockListener.On("Close").Return(errors.New("SOMETHING_WRONG"))
				mockNet := &mockNet{}
				mockNet.On("Listen", "tcp", ":0").Return(mockListener, nil)
				return &mocks{mockNet, mockListener}
			}(),
			0,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			l := &Listener{
				listen: tt.mocks.net.Listen,
			}

			got, err := l.RandomUnusedTCPPort()

			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			tt.mocks.net.AssertExpectations(t)
			tt.mocks.listener.AssertExpectations(t)
		})
	}
}
