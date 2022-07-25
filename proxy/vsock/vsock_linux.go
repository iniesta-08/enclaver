package vsock

import (
	"github.com/mdlayher/vsock"
	"net"
)

const (
	// parentContextID is a fixed constant value of 3 defined by AWS
	parentContextID uint32 = 3
)

func DialParent(port uint32) (net.Conn, error) {
	return vsock.Dial(parentContextID, port, nil)
}

func DialEnclave(contextID uint32, port uint32) (net.Conn, error) {
	// TODO: What to do here?
	return vsock.Dial(contextID, port, nil)
}

func Listen(port uint32) (net.Listener, error) {
	return vsock.Listen(port, nil)
}