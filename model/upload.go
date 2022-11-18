package model

type EdgeUploadResponse struct {
	Destination string `json:"destination"`
	File        string `json:"file"`
	Size        string `json:"size"`
	UploadTime  string `json:"upload_time"`
}

type FileUpload struct {
	Arch    string `json:"arch" binding:"required"`
	Version string `json:"version" binding:"required"`
	File    string `json:"file"`
}
