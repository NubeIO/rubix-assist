package systemd

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/utils/command"
	"strings"
)

//
//action: "stop/start/restart/enable/disable"
//service: "POINT_SERVER"

// SystemctlStatus Run `systemctl status is-active`
func SystemctlStatus(services []string) (statuses map[string]string, success bool) {
	statuses = make(map[string]string)

	args := []string{"systemctl", "is-active"}
	args = append(args, services...)

	output, _ := command.SudoRun(args...)
	for i, status := range strings.Split(output, "\n") {
		statuses[services[i]] = status
	}

	return statuses, true
}

// SystemctlStart Run `systemctl start [service]`
func SystemctlStart(service string) (message string, success bool) {
	if output, err := command.Run("sudo systemctl", "start", service); err == nil {
		return output, true
	} else {
		return fmt.Sprintf("Failed to start service: %s", service), false
	}
}

// SystemctlStop Run `systemctl stop [service]`
func SystemctlStop(service string) (message string, success bool) {
	if output, err := command.Run("sudo systemctl", "stop", service); err == nil {
		return output, true
	} else {
		return fmt.Sprintf("Failed to stop service: %s", service), false
	}
}

// SystemctlRestart Run `systemctl restart [service]`
func SystemctlRestart(service string) (message string, success bool) {
	if output, err := command.Run("sudo systemctl", "restart", service); err == nil {
		return output, true
	} else {
		return fmt.Sprintf("Failed to restart service: %s", service), false
	}
}
