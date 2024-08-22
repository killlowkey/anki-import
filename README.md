## 有道翻译单词发音接口
默认的是美音：http://dict.youdao.com/dictvoice?audio=单词

等号后面接单词，例如：http://dict.youdao.com/dictvoice?audio=the

英音和美音（需要进行 url 编码）：
1. 英音：http://dict.youdao.com/dictvoice?type=1&audio=
2. 美音：http://dict.youdao.com/dictvoice?type=0&audio=

## 有道翻译单词解释查询接口
http://dict.youdao.com/suggest?num=1&doctype=json&q=单词

等号后面接单词，例如：http://dict.youdao.com/suggest?num=1&doctype=json&q=the

## 有道翻译句子接口
http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i=要翻译的名子或短语

等号后面接句子或短语，例如：
http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i=English-speaking world

如果是链接，将句子的空格替换成%20，例如：
http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i=English-speaking%20world

## microsoft tts
1. https://github.com/rany2/edge-tts
   ```text
   word.csv 单词
   字段名	类型	说明	示例
   vc_id	string	单词id	57067c89a172044907c6698e
   vc_vocabulary	string	单词	superspecies
   vc_phonetic_uk	string	uk英音音标	[su:pərsˈpi:ʃi:z]
   vc_phonetic_us	string	us美音音标	[supɚsˈpiʃiz]
   vc_frequency	float	词频	0.000000
   vc_difficulty	int	难度	1
   vc_acknowledge_rate	float	认识率	0.664122
   
   word_translation.csv 单词及其中文翻译
   字段名	类型	说明	示例
   word	string	单词	brain
   translation	string	单词的中文翻译	n.脑,头脑
    ```
2. https://github.com/surfaceyu/edge-tts-go


## anki connect plugin
1. https://ankiweb.net/shared/info/2055492159

## 本地词库
1. https://github.com/LinXueyuanStdio/DictionaryData
2. https://github.com/skywind3000/ECDICT

## 卡片模板

### 字段
在修改卡片正面和反面下方可以添加和修改字段
```text
word            // 单词
ipa_uk          // 英式音标
ipa_us           // 美式音标
ipa_audio       // 音频
definition_cn    // 翻译
source_name1    // 来源1
source_content1 // 来源1内容
source_translate1  // 来源1翻译
source_name2    // 来源2
source_content2 // 来源2内容
source_translate2  // 来源2翻译
examples1_en     // 例子2内容
examples1_cn    // 例子2翻译
examples2_en     // 例子2内容
examples2_cn     // 例子2翻译
```

### 正面
```html
<div class="section">
    <div id="front" class="items">
        <span id="word">{{word}}</span>
        <span id="ipa">{{ipa_uk}}</span>
        <span id="ipa">{{ipa_us}}</span>
        {{ipa_audio}}
    </div>
</div>

<div class="section">
    <div class="items content">{{source_content1}}</div>
    <div class="items name">{{source_name1}}</div>
</div>

<div class="section">
    <div class="items content">{{source_content2}}</div>
    <div class="items name">{{source_name2}}</div>
</div>
```

### 背面
```html
<div class="section">
    <div id="front" class="items">
        <span id="word">{{word}}</span>
        <span id="ipa">{{ipa_uk}}</span>
        <span id="ipa">{{ipa_us}}</span>
        {{ipa_audio}}
    </div>
    <div id="back" class="items">{{definition_cn}}</div>
</div>

<div class="section">
    <div class="items content">{{source_content1}}</div>
    <div class="items content">{{source_translate1}}</div>
    <div class="items name">{{source_name1}}</div>
</div>

<div class="section">
    <div class="items content">{{source_content2}}</div>
    <div class="items content">{{source_translate2}}</div>
    <div class="items name">{{source_name2}}</div>
</div>

<div class="section">
    <div class="items content">{{examples1_en}}</div>
    <div class="items content">{{examples1_cn}}</div>
</div>

<div class="section">
    <div class="items content">{{examples2_en}}</div>
    <div class="items content">{{examples2_cn}}</div>
</div>
```

### style
```css
</style>

<style>

.card {
  margin: 12px;
  text-align: left;
  background-color: #fdfdfd;
}

.section {
  color: #414141;
  background-color: #fefffe;
  font-family: Arial, sans-serif;
  font-size: 16px;
  box-shadow: 1px 1px 5px 0px rgba(0, 0, 0, 0.3),0 0px 0px 1px rgba(0, 0, 0, 0);
  border-radius: 5px;
  margin: 8px 0;
}

.items {
  margin: 0 12px;
  padding: 4px 0;
}

#front, #back {
  line-height: 1.4em;
}

#front {
  font-size: 1.4em;
  font-weight: bold;
  text-align: left;
}

#back {
  font-size: 1em;
  font-weight: normal;
  text-align: left;
}

#ipa {
  font-size: 0.8em;
  font-weight: normal;
  font-style:italic;
}

.content {
  font-size: 1em
}

.name {
  font-size: 0.8em;
  text-align:right;
  font-style:italic;
}

div:empty {
	display: none
}

a {
  text-decoration:none;
  color:inherit
}


</style>

<style>
```