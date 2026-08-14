package types

type MPVCommand struct {
	Command []interface{} `json:"command"`
}

type ServerUnix_StatusResponse struct {
	Data      string `json:"data"`
	RequestID int    `json:"request_id"`
	Error     string `json:"error"`
}
