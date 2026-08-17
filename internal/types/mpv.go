package types

type MediaItem struct {
	Current  bool   `json:"current,omitempty"`
	Filename string `json:"filename"`
	ID       int    `json:"id"`
	Playing  bool   `json:"playing,omitempty"`
}

type Response struct {
	Data      []MediaItem `json:"data"`
	Error     string      `json:"error"`
	RequestID int         `json:"request_id"`
}
