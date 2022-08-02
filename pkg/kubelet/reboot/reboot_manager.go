package reboot

import (
	"k8s.io/kubernetes/pkg/kubelet/nodeshutdown"
)

// TODO: think about more elegant approach than SetNodeShutdownManager()

type Manager interface {
	SetNodeShutdownManager(mgr nodeshutdown.Manager)
	Run() error
}

var _ Manager = (*managerStub)(nil)

type managerStub struct { }

func (h *managerStub) SetNodeShutdownManager(mgr nodeshutdown.Manager) { }

func (h *managerStub) Run() error {
	return nil
}