//go:build !linux
// +build !linux

package reboot


func NewManager() Manager {
	return &managerStub{}
}
