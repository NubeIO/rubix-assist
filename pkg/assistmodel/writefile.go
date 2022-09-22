package assistmodel

type WriteFile struct {
	FilePath     string      `json:"path"`
	Body         interface{} `json:"body"`
	BodyAsString string      `json:"body_as_string"`
}
