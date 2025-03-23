package transport

type HandShakeHandler func(any) error

// NOPHandShakeHandler No operation, test the connection, like noop
func NOPHandShakeHandler(any) error {
	return nil
}
