package tts

import (
	"errors"
	"github.com/surfaceyu/edge-tts-go/edgeTTS"
)

var displayShortMap = map[string]string{
	"晓晓":        "zh-CN-XiaoxiaoNeural",
	"晓伊":        "zh-CN-XiaoyiNeural",
	"云健":        "zh-CN-YunjianNeural",
	"云希":        "zh-CN-YunxiNeural",
	"云夏":        "zh-CN-YunxiaNeural",
	"云扬":        "zh-CN-YunyangNeural",
	"Aria":        "en-US-AriaNeural",
	"Guy":         "en-US-GuyNeural",
	"Jenny":       "en-US-JennyNeural",
	"Amber":       "en-US-AmberNeural",
	"Ana":         "en-US-AnaNeural",
	"Christopher": "en-US-ChristopherNeural",
	"Eric":        "en-US-EricNeural",
	"Elizabeth":   "en-US-ElizabethNeural",
	"Davis":       "en-US-DavisNeural",
	"Jane":        "en-US-JaneNeural",
	"Michelle":    "en-US-MichelleNeural",
	"Tony":        "en-US-TonyNeural",
	"Ashley":      "en-GB-AshleyNeural",
	"Ryan":        "en-GB-RyanNeural",
	"Libby":       "en-GB-LibbyNeural",
	"Thomas":      "en-GB-ThomasNeural",
	"Arthur":      "en-AU-ArthurNeural",
	"Carly":       "en-AU-CarlyNeural",
	"Natasha":     "en-IN-NatashaNeural",
	"Ravi":        "en-IN-RaviNeural",
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Speak(speaker, text, output string) error {
	tts := edgeTTS.NewTTS(edgeTTS.Args{
		Voice:      "",
		WriteMedia: output,
	})

	if s, ok := displayShortMap[speaker]; !ok {
		return errors.New("speaker not found in displayShortMap")
	} else {
		speaker = s
	}

	tts.AddTextWithVoice(text, speaker)
	tts.Speak()
	return nil
}
