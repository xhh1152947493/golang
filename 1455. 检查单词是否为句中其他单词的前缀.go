package main

import (
	"fmt"
	"strings"
)

/*
给你一个字符串 sentence 作为句子并指定检索词为 searchWord ，其中句子由若干用 单个空格 分隔的单词组成。

请你检查检索词 searchWord 是否为句子 sentence 中任意单词的前缀。

如果 searchWord 是某一个单词的前缀，则返回句子 sentence 中该单词所对应的下标（下标从 1 开始）。
如果 searchWord 是多个单词的前缀，则返回匹配的第一个单词的下标（最小下标）。
如果 searchWord 不是任何单词的前缀，则返回 -1 。
字符串 S 的 「前缀」是 S 的任何前导连续子字符串。

1 <= sentence.length <= 100
1 <= searchWord.length <= 10
sentence 由小写英文字母和空格组成。
searchWord 由小写英文字母组成。
前缀就是紧密附着于词根的语素，中间不能插入其它成分，并且它的位置是固定的——-位于词根之前。（引用自 前缀_百度百科 ）
*/

func isPrefixOfWord(sentence string, searchWord string) int {
	sentenceArr := strings.Split(sentence, " ") // 分割字符串
	for i, word := range sentenceArr {
		flag := strings.HasPrefix(word, searchWord) // 以前缀开始
		if flag {
			return i + 1
		}
	}
	return -1
}

func main() {
	var test = "hellohello hellohellohello"
	var test1 = "ell"
	fmt.Println(isPrefixOfWord(test, test1))
}
