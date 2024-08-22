package anki

type Request struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  any    `json:"params"`
}

type AddNoteResp struct {
	Result []int64 `json:"result"`
	Error  string  `json:"error"`
}
