package model

import "github.com/NubeIO/lib-systemctl-go/systemctl"

type SystemCtlBody struct {
	AppName      string   `json:"app_name"`      // "flow-framework"
	ServiceName  string   `json:"service_name"`  // "nubeio-flow-framework.service"
	Action       string   `json:"action"`        // start, stop, restart, enable, disable
	AppNames     []string `json:"app_names"`     // ["nubeio-rubix-edge.service", "nubeio-flow-framework.service"]
	ServiceNames []string `json:"service_names"` // ["rubix-edge", "flow-framework"]
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
