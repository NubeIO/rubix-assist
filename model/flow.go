package model

type FFNetwork struct {
	NetworkName string   `json:"network_name"`
	PluginName  string   `json:"plugin_name"`
	DeviceName  string   `json:"device_name"`
	Points      []string `json:"points"`
}

type FFFlowNetwork struct {
	Name                   string   `json:"name"`
	FetchHistories         bool     `json:"fetch_histories"`
	FetchHistFrequency     int      `json:"fetch_hist_frequency"`
	DeleteHistoriesOnFetch bool     `json:"delete_histories_on_fetch"`
	IsMasterSlave          bool     `json:"is_master_slave"`
	IsMqtt                 bool     `json:"is_mqtt"`
	FlowHttps              bool     `json:"flow_https"`
	FlowIp                 string   `json:"flow_ip"`
	FlowPort               int      `json:"flow_port"`
	FlowUsername           string   `json:"flow_username"`
	FlowPassword           string   `json:"flow_password"`
	FlowToken              string   `json:"flow_token"`
	StreamName             string   `json:"stream_name"`
	ProducersNames         []string `json:"producer_names"`
}
