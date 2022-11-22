package model

type Plugin struct {
	Name      string `json:"name"`
	Arch      string `json:"arch"`
	Version   string `json:"version,omitempty"`
	Extension string `json:"extension"`
}

type PluginUpload struct {
	Name                 string `json:"name"`
	Arch                 string `json:"arch"`
	Version              string `json:"version,omitempty"`
	Extension            string `json:"extension"`
	ClearBeforeUploading bool   `json:"clear_before_uploading"`
}
