package transport

import "testing"

func TestNewTCP(t *testing.T) {
	address := "8080"
	tcp := NewTCP(address)

	if tcp.address != address {
		t.Error("NewTCP address mismatch")
	}
	err := tcp.Listen()
	if err != nil {
		t.Error(err)
	}
}
