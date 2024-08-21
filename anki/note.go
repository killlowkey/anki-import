package anki

// DuplicateScopeOptions 用于定义在添加笔记时检查重复的范围和条件。
type DuplicateScopeOptions struct {
	DeckName       string `json:"deckName"`       // 要检查重复的牌组名称
	CheckChildren  bool   `json:"checkChildren"`  // 是否检查子牌组
	CheckAllModels bool   `json:"checkAllModels"` // 是否检查所有模型
}

// Options 定义了在添加笔记时的一些选项。
type Options struct {
	AllowDuplicate        bool                  `json:"allowDuplicate"`        // 是否允许重复的笔记
	DuplicateScope        string                `json:"duplicateScope"`        // 重复检查的范围，可以是"deck"（仅当前牌组）或"collection"（整个集合）
	DuplicateScopeOptions DuplicateScopeOptions `json:"duplicateScopeOptions"` // 详细的重复检查选项
}

// Fields 用于定义笔记的字段内容，比如前面和后面的内容。
type Fields struct {
	Front string `json:"Front"` // 笔记的正面内容
	Back  string `json:"Back"`  // 笔记的背面内容
}

// Media 定义了与笔记关联的媒体文件，如音频、视频或图片。
type Media struct {
	URL      string   `json:"url"`      // 媒体文件的 URL 地址
	Filename string   `json:"filename"` // 保存到 Anki 中的文件名
	SkipHash string   `json:"skipHash"` // 用于跳过重复媒体文件的哈希值
	Fields   []string `json:"fields"`   // 关联的笔记字段，媒体文件将显示在这些字段上
}

// Note 表示要添加到 Anki 中的一条笔记。
type Note struct {
	DeckName  string            `json:"deckName"`  // 笔记将添加到的牌组名称
	ModelName string            `json:"modelName"` // 使用的笔记模型名称
	Fields    map[string]string `json:"fields"`    // 笔记的字段内容，包括正面和背面
	Options   Options           `json:"options"`   // 笔记的选项，包括重复检查设置
	Tags      []string          `json:"tags"`      // 笔记的标签列表
	Audio     []Media           `json:"audio"`     // 与笔记关联的音频文件
	Video     []Media           `json:"video"`     // 与笔记关联的视频文件
	Picture   []Media           `json:"picture"`   // 与笔记关联的图片文件
}
