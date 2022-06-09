package assitcli

type Ping struct {
	Version                string `json:"version"`
	UpTimeDate             string `json:"up_time_date"`
	UpMin                  string `json:"up_min"`
	UpHour                 string `json:"up_hour"`
	DeploymentMode         string `json:"deployment_mode"`
	PublicIp               string `json:"public_ip"`
	MqttRestBridgeListener struct {
		Enabled bool `json:"enabled"`
		Master  bool `json:"master"`
		Status  bool `json:"status"`
	} `json:"mqtt_rest_bridge_listener"`
}
