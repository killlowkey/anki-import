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

type FindNotesResp struct {
	Result []int64 `json:"result"`
	Error  string  `json:"error"`
}

type NotesInfoResp struct {
	Result []NoteInfo `json:"result"`
	Error  string     `json:"error"`
}

type DeleteNoteResp struct {
	Result any    `json:"result"`
	Error  string `json:"error"`
}

type UpdateNoteReq struct {
	Id     int64             `json:"id"`     // note id
	Fields map[string]string `json:"fields"` // note 字段
	Tags   []string          `json:"tags"`   // 标签
}

type UpdateNoteResp struct {
	Result any    `json:"result"`
	Error  string `json:"error"`
}
