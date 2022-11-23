package amodel

import "github.com/NubeIO/lib-systemctl-go/systemctl"

type Action string

const (
	Enable  Action = "enable"
	Disable Action = "disable"
	Start   Action = "start"
	Stop    Action = "stop"
	Restart Action = "restart"
)

type SystemCtlBody struct {
	AppName string `json:"app_name"` // "flow-framework"
	Action  Action `json:"action"`   // start, stop, restart, enable, disable
}

type SystemResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AppSystemState struct {
	ServiceName            string                  `json:"service_name,omitempty"`
	AppName                string                  `json:"app_name,omitempty"`
	State                  systemctl.UnitFileState `json:"state,omitempty"`        // enabled, disabled
	ActiveState            systemctl.ActiveState   `json:"active_state,omitempty"` // active, inactive
	SubState               systemctl.SubState      `json:"sub_state,omitempty"`    // running, dead
	ActiveEnterTimestamp   string                  `json:"active_enter_timestamp,omitempty"`
	InactiveEnterTimestamp string                  `json:"inactive_enter_timestamp,omitempty"`
	Restarts               string                  `json:"restarts,omitempty"` // number of restart
	IsInstalled            bool                    `json:"is_installed,omitempty"`
}

type MassSystemResponse struct {
	AppName     string `json:"app_name"`
	ServiceName string `json:"service_name"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
}
