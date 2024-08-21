package tts

import (
	"fmt"
	"github.com/surfaceyu/edge-tts-go/edgeTTS"
	"strings"
	"testing"
	"time"
)

func Test_Speak(t *testing.T) {
	// edgeTTS.NewTTS(args).AddText(args.Text, args.Voice, args.Rate, args.Volume).Speak()
	contents := []string{
		"[[Aria]] Microsoft's TTS voices are all designed to sound natural, but some voices are particularly noted for their human-like quality. Generally, the \"Neural\" voices in the Microsoft TTS lineup are the most advanced in terms of naturalness, as they use neural networks to generate speech. Among these, the following voices are often considered to sound especially natural:",
	}
	args := edgeTTS.Args{
		Voice:      "",
		WriteMedia: "./sample.mp3",
	}
	start := time.Now()
	tts := edgeTTS.NewTTS(args)
	for _, v := range contents {
		speaker, text := splitSpeaker(v)
		tts.AddTextWithVoice(text, speaker)
	}
	tts.Speak()
	fmt.Printf("程序运行时间: %s", time.Since(start))
}

// get short name by display name
func SpeakerShortName(display string) string {
	display, ok := displayShortMap[display]
	if ok {
		return display
	}
	return ""
}

func splitSpeaker(content string) (string, string) {
	startIndex := strings.Index(content, "[[")
	endIndex := strings.Index(content, "]]")
	if startIndex == -1 || endIndex == -1 {
		return "", content
	}
	speaker := SpeakerShortName(content[startIndex+2 : endIndex])
	annotation := content[endIndex+2:]
	return speaker, annotation
}
