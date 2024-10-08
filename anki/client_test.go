package anki

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddNote1(t *testing.T) {
	note := Note{
		DeckName:  "哈利波特与魔法石",
		ModelName: "问答题",
		Fields: map[string]string{
			"正面": "front content",
			"背面": "back content",
		},
		Options: Options{
			AllowDuplicate: false,
			DuplicateScope: "deck",
			DuplicateScopeOptions: DuplicateScopeOptions{
				DeckName:       "Default",
				CheckChildren:  false,
				CheckAllModels: false,
			},
		},
		Tags: []string{"哈利波特与魔法石"},
		Audio: []Media{
			{
				URL:      "https://assets.languagepod101.com/dictionary/japanese/audiomp3.php?kanji=猫&kana=ねこ",
				Filename: "yomichan_ねこ_猫.mp3",
				SkipHash: "7e2c2f954ef6051373ba916f000168dc",
				Fields:   []string{"正面"},
			},
		},
		Video: []Media{
			{
				URL:      "https://cdn.videvo.net/videvo_files/video/free/2015-06/small_watermarked/Contador_Glam_preview.mp4",
				Filename: "countdown.mp4",
				SkipHash: "4117e8aab0d37534d9c8eac362388bbe",
				Fields:   []string{"背面"},
			},
		},
		Picture: []Media{
			{
				URL:      "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/A_black_cat_named_Tilly.jpg/220px-A_black_cat_named_Tilly.jpg",
				Filename: "black_cat.jpg",
				SkipHash: "8d6e4646dfae812bf39651b59d7429ce",
				Fields:   []string{"背面"},
			},
		},
	}

	client := NewClient("http://localhost:8765", WithDebug())
	client.AddNote(note)
}

func Test_AddNote2(t *testing.T) {
	note := Note{
		DeckName:  "哈利波特与魔法石",
		ModelName: "english-word",
		Fields: map[string]string{
			"word":              "example",                      // 单词
			"ipa_uk":            "/ɪɡˈzɑːm.pəl/",                // 英式音标
			"ipa_us":            "/ɪɡˈzæm.pəl/",                 // 美式音标
			"ipa_audio":         "",                             // 音频
			"definition_cn":     "n. 例子；实例<br>n. 例子；实例",         // 翻译
			"source_name1":      "example1",                     // 来源1
			"source_content1":   "This is an example sentence.", // 来源1内容
			"source_translate1": "这是一个例句。",                      // 来源1翻译
			"source_name2":      "",                             // 来源2
			"source_content2":   "",                             // 来源2内容
			"source_translate2": "",                             // 来源2翻译
			"examples1_en":      "",                             // 例子2内容
			"examples1_cn":      "",                             // 例子2翻译
			"examples2_en":      "",                             // 例子2内容
			"examples2_cn":      "",                             // 例子2翻译
		},
		Tags: []string{"哈利波特与魔法石"},
		Audio: []Media{
			{
				URL:      "http://dict.youdao.com/dictvoice?type=0&audio=example",
				Filename: "example.mp3",
				SkipHash: "7e2c2f954ef6051373ba916f000168dc",
				Fields:   []string{"ipa_audio"}, // 关联音频
			},
		},
	}

	client := NewClient("http://localhost:8765", WithDebug())
	client.AddNote(note)
}

func TestClient_All(t *testing.T) {
	client := NewClient("http://localhost:8765", WithDebug())

	// 添加
	note := Note{
		DeckName:  "哈利波特与魔法石",
		ModelName: "english-word",
		Fields: map[string]string{
			"word":              "test-test-test",               // 单词
			"ipa_uk":            "/ɪɡˈzɑːm.pəl/",                // 英式音标
			"ipa_us":            "/ɪɡˈzæm.pəl/",                 // 美式音标
			"ipa_audio":         "",                             // 音频
			"definition_cn":     "n. 例子；实例<br>n. 例子；实例",         // 翻译
			"source_name1":      "example1",                     // 来源1
			"source_content1":   "This is an example sentence.", // 来源1内容
			"source_translate1": "这是一个例句。",                      // 来源1翻译
			"source_name2":      "",                             // 来源2
			"source_content2":   "",                             // 来源2内容
			"source_translate2": "",                             // 来源2翻译
			"examples1_en":      "",                             // 例子2内容
			"examples1_cn":      "",                             // 例子2翻译
			"examples2_en":      "",                             // 例子2内容
			"examples2_cn":      "",                             // 例子2翻译
		},
		Tags: []string{"哈利波特与魔法石"},
		Audio: []Media{
			{
				URL:      "http://dict.youdao.com/dictvoice?type=0&audio=example",
				Filename: "example.mp3",
				SkipHash: "7e2c2f954ef6051373ba916f000168dc",
				Fields:   []string{"ipa_audio"}, // 关联音频
			},
		},
	}
	noteId, err := client.AddNote(note)
	assert.NoError(t, err)
	assert.NotEmpty(t, noteId)

	// 获取 note 信息
	noteInfo, err := client.NoteInfo(noteId)
	assert.NoError(t, err)
	assert.NotEmpty(t, noteInfo)
	assert.Equal(t, "test-test-test", noteInfo.Fields["word"].Value)
	assert.Equal(t, 1, len(noteInfo.Tags))
	assert.Equal(t, "哈利波特与魔法石", noteInfo.Tags[0])

	// 更新 note
	err = client.UpdateNote(UpdateNoteReq{
		Id: noteId,
		Fields: map[string]string{
			"word": "update-test-test-test",
		},
		Tags: []string{"哈利波特与魔法石1"},
	})
	assert.NoError(t, err)

	// 更新是否成功
	noteInfo, err = client.NoteInfo(noteId)
	assert.NoError(t, err)
	assert.Equal(t, "update-test-test-test", noteInfo.Fields["word"].Value)
	assert.Equal(t, []string{"哈利波特与魔法石1"}, noteInfo.Tags)

	// 删除 note
	err = client.DeleteNote([]int64{noteId})
	assert.NoError(t, err)

	// 是否存在
	_, err = client.NoteInfo(noteId)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "note id is empty")
}

func TestClient_FindNotes(t *testing.T) {
	client := NewClient("http://localhost:8765", WithDebug())
	noteIds, err := client.FindNotes("deck:哈利波特与魔法石")
	assert.NoError(t, err)
	assert.NotEmpty(t, noteIds)

	noteInfos, err := client.NotesInfo(noteIds)
	assert.NoError(t, err)
	assert.NotEmpty(t, noteInfos)
}
