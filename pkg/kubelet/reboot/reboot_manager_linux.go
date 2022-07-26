package reboot

import (
	"os/exec"
	"syscall"

	"k8s.io/kubernetes/pkg/kubelet/nodeshutdown"
)

func NewManager() Manager {
	return &manager{}
}

var _ Manager = (*manager)(nil)

type manager struct {
	nodeShutdownManager nodeshutdown.Manager
}

func (h *manager) SetNodeShutdownManager(mgr nodeshutdown.Manager) {
	h.nodeShutdownManager = mgr
}

func (h *manager) Run() error {
	if err := h.nodeShutdownManager.TriggerShutdownProcedure(); err != nil {
		return err
	}

	if err := exec.Command("sync").Run(); err != nil {
		return err
	}

	return syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
}
