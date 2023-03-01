package interfaces

type CreateSnapshot struct {
	Description string `json:"description"`
}

type RestoreSnapshot struct {
	File        string `json:"file"`
	Description string `json:"description"`
}
