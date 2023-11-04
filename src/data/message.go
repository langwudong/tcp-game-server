package data

type Message struct {
	Type    string `json:"type"`
	Message string `json:"message,omitempty"`
}
