package model

type Message struct {
	Message interface{} `json:"message"`
}

type FoundMessage struct {
	Found bool `json:"found"`
}
