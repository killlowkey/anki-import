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

func (c *Client) postRequest(body any) (string, error) {
	resp, err := c.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("/")

	if err != nil {
		return "", fmt.Errorf("error occurred: %v", err)
	}

	return resp.String(), nil
}

func (c *Client) AddNotes(notes []Note) (string, error) {
	data := RequestData{
		Action:  "addNotes",
		Version: 6,
		Params: map[string]any{
			"notes": notes,
		},
	}

	return c.postRequest(data)
}

func (c *Client) AddNote(note Note) (string, error) {
	return c.AddNotes([]Note{note})
}
