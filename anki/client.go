package anki

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// Client
// https://ankiweb.net/shared/info/2055492159
// https://foosoft.net/projects/anki-connect/
type Client struct {
	httpClient *resty.Client
}

type Option func(*Client)

func WithDebug() Option {
	return func(client *Client) {
		client.httpClient.SetDebug(true)
	}
}

func NewClient(baseUrl string, options ...Option) *Client {
	client := &Client{
		httpClient: resty.New().SetBaseURL(baseUrl),
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func (c *Client) SetDebug(debug bool) {
	c.httpClient.SetDebug(debug)
}

func (c *Client) postRequest(body any, resp any) error {
	resp, err := c.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&resp).
		Post("/")

	if err != nil {
		return fmt.Errorf("error occurred: %v", err)
	}

	return nil
}

// AddNotes 批量添加笔记，返回笔记 id
func (c *Client) AddNotes(notes []Note) ([]int64, error) {
	data := Request{
		Action:  "addNotes",
		Version: 6,
		Params: map[string]any{
			"notes": notes,
		},
	}

	var addNotesResp AddNoteResp
	if err := c.postRequest(data, &addNotesResp); err != nil {
		return nil, err
	}

	if addNotesResp.Error != "" {
		return nil, fmt.Errorf("error occurred: %v", addNotesResp.Error)
	}

	return addNotesResp.Result, nil
}

// AddNote 添加笔记，返回笔记 id
func (c *Client) AddNote(note Note) (int64, error) {
	if notes, err := c.AddNotes([]Note{note}); err != nil {
		return -1, err
	} else if len(notes) <= 0 {
		return -1, fmt.Errorf("note id is empty")
	} else {
		return notes[0], nil
	}
}
