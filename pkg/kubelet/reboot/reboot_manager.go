package reboot

import (
	"k8s.io/kubernetes/pkg/kubelet/nodeshutdown"
)

// TODO: think about more elegant approach than SetNodeShutdownManager()

type Manager interface {
	SetNodeShutdownManager(mgr nodeshutdown.Manager)
	Run() error
}
