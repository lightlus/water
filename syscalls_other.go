// +build !linux,!darwin,!windows

package water

import "errors"

func openDev(config Config) (*Interface, error) {
	return nil, errors.New("not implemented on this platform")
}

func updateTunNetwork(ifce *Interface, network string) error {
	return nil
}
