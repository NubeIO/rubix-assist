package assistmodel

type EdgeConfig struct {
	AppName      string      `json:"app_name,omitempty"`
	Body         interface{} `json:"body"`                  // used when writing JSON, YML data
	BodyAsString string      `json:"body_as_string"`        // used when writing string data
	ConfigType   string      `json:"config_type,omitempty"` // config.yml
}

type EdgeConfigResponse struct {
	FilePath string `json:"file_path,omitempty"`
	Data     []byte `json:"data"`
}
