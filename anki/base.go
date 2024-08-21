package anki

type RequestData struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  any    `json:"params"`
}
