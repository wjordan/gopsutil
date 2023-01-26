//go:build !linux
// +build !linux

package net

func NetInterfaces() ([]net.Interface, error) {
	return net.Interfaces()
}
