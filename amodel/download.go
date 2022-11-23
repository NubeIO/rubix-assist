package amodel

type EdgeDownloadResponse struct {
	FileName    string `json:"file,omitempty"`
	Path        string `json:"path,omitempty"`
	Destination string `json:"destination"`
}
