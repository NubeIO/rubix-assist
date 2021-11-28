package model

type LocalstorageFlowNetwork struct {
	FlowIp       string `json:"flow_ip"`
	FlowPort     int    `json:"flow_port"`
	FlowHttps    bool   `json:"flow_https"`
	FlowUsername string `json:"flow_username"`
	FlowPassword string `json:"flow_password"`
}
