package reboot

import (
	"os/exec"

	"k8s.io/kubernetes/pkg/kubelet/nodeshutdown"

	"github.com/coreos/go-systemd/v22/login1"
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
	conn, err := login1.New()
	if err != nil {
		return err
	}

	if err := h.nodeShutdownManager.TriggerShutdownProcedure(); err != nil {
		return err
	}

	if err := exec.Command("sync").Run(); err != nil {
		return err
	}

	// return syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)

	conn.Reboot(false)
	return nil
}
