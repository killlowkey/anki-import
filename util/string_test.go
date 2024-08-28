package util

import (
	"reflect"
	"testing"
)

// 测试 ExtractTranslations 函数
func TestExtractTranslations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "Basic verbs and nouns",
			input: "v. 俯冲；突然袭击；（尤指鸟）猛扑；（非正式）猛地抓起; 和 n. 猛扑；突然袭击",
			expected: []string{
				"v. 俯冲；突然袭击；（尤指鸟）猛扑；（非正式）猛地抓起",
				"n. 猛扑；突然袭击",
			},
		},
		{
			name:  "Multiple parts of speech",
			input: "v. 跑；跳；走;n. 跑步；跳跃;adj. 快速的；突然的",
			expected: []string{
				"v. 跑；跳；走",
				"n. 跑步；跳跃",
				"adj. 快速的；突然的",
			},
		},
		{
			name:     "Single part of speech",
			input:    "adv. 慢慢地",
			expected: []string{"adv. 慢慢地"},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:  "Input with extra spaces",
			input: "   v.   缓慢地  ；  静静地   ; n. 宁静； 平静 ",
			expected: []string{
				"v. 缓慢地  ；  静静地",
				"n. 宁静； 平静",
			},
		},
		{
			name:  "Complex input with multiple delimiters",
			input: "v. 俯冲, 突然袭击; n. 猛扑，突然袭击; adj. 快速的；突然的",
			expected: []string{
				"v. 俯冲, 突然袭击",
				"n. 猛扑，突然袭击",
				"adj. 快速的；突然的",
			},
		},
		{
			name:  "Only part of speech tags with no definitions",
			input: "v.; n.; adj.",
			expected: []string{
				"v. ",
				"n. ",
				"adj. ",
			},
		},
		{
			name:     "No valid part of speech tags",
			input:    "This is a test without valid tags.",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ExtractTranslations(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("ExtractTranslations(%q) = %v, expected %v", tt.input, actual, tt.expected)
			}
		})
	}
}
