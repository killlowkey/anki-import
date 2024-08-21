package youdao

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *resty.Client
}

type Option func(client *Client)

func WithDebug() Option {
	return func(client *Client) {
		client.httpClient.SetDebug(true)
	}
}

func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient: resty.New(),
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// TranslateWord 翻译单词
func (c *Client) TranslateWord(word string) (explain string, entry string, err error) {
	type TranslateResp struct {
		Result struct {
			Msg  string `json:"msg"`
			Code int    `json:"code"`
		} `json:"result"`
		Data struct {
			Entries []struct {
				Explain string `json:"explain"`
				Entry   string `json:"entry"`
			} `json:"entries"`
			Query    string `json:"query"`
			Language string `json:"language"`
			Type     string `json:"type"`
		} `json:"data"`
	}

	var resp TranslateResp
	_, err = c.httpClient.R().
		SetResult(&resp).
		Get("https://dict.youdao.com/suggest?num=1&doctype=json&q=" + url.QueryEscape(word))
	if err != nil {
		return
	}

	if resp.Result.Code != http.StatusOK {
		err = fmt.Errorf("code=%d, msg=%s", resp.Result.Code, resp.Result.Msg)
		return
	}

	if len(resp.Data.Entries) <= 0 {
		err = fmt.Errorf("entries is empty")
		return
	}

	explain = resp.Data.Entries[0].Explain
	entry = resp.Data.Entries[0].Entry
	return
}

// AudioUk 单词英音
func (c *Client) AudioUk(word string) string {
	return "http://dict.youdao.com/dictvoice?type=1&audio=" + word
}

// AudioUS 单词美音
func (c *Client) AudioUS(word string) string {
	return "http://dict.youdao.com/dictvoice?type=0&audio=" + word
}
