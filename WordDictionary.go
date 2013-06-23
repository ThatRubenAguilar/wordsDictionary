package words

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

type MapNode struct {
	nodeMap   map[rune]*MapNode
	endOfWord bool
}

type WordDictionary struct {
	wordDict *MapNode
}

type WordInfo struct {
	Word     string
	IsPrefix bool
	IsWord   bool
}

func (SrcWordDict *WordDictionary) Copy() (DestWordDict *WordDictionary) {
	DestWordDict = new(WordDictionary)
	if SrcWordDict.wordDict != nil {
		DestWordDict.wordDict = new(MapNode)
		var SrcNode *MapNode = SrcWordDict.wordDict
		var DestNode *MapNode = DestWordDict.wordDict
		DestNode.endOfWord = SrcNode.endOfWord
		copyRecurse(SrcNode, DestNode)
	}
	return
}
func copyRecurse(SrcNode *MapNode, DestNode *MapNode) {
	DestNode.endOfWord = SrcNode.endOfWord
	if SrcNode.nodeMap != nil {
		DestNode.nodeMap = make(map[rune]*MapNode)
		for srcKey, srcValue := range SrcNode.nodeMap {
			if srcValue != nil {
				DestNode.nodeMap[srcKey] = new(MapNode)
				DestNode.nodeMap[srcKey].endOfWord = srcValue.endOfWord
				copyRecurse(srcValue, DestNode.nodeMap[srcKey])
			} else {
				DestNode.nodeMap[srcKey] = nil
			}
		}
	}
	return
}

func (WordDict *WordDictionary) Lookup(Word string) (wordInfo WordInfo, err error) {
	wordInfo.Word = Word
	Word = strings.ToLower(Word)

	var CurrentNode *MapNode = WordDict.wordDict
	var Runes []rune = []rune(Word)
	for _, Rune := range Runes {
		if CurrentNode.nodeMap == nil {
			wordInfo.IsPrefix = false
			wordInfo.IsWord = false
			return
		}
		_, exist := CurrentNode.nodeMap[Rune]
		if !exist {
			wordInfo.IsPrefix = false
			wordInfo.IsWord = false
			return
		}
		CurrentNode = CurrentNode.nodeMap[Rune]
	}
	wordInfo.IsWord = CurrentNode.endOfWord
	wordInfo.IsPrefix = true
	return
}

func (WordDict *WordDictionary) AddWords(Words []string) (err error) {

	if WordDict.wordDict == nil {
		WordDict.wordDict = new(MapNode)
		WordDict.wordDict.endOfWord = false
	}
	// For every char of every string hash it in the word dict
	for _, Word := range Words {
		Word = strings.ToLower(Word)
		var CurrentNode *MapNode = WordDict.wordDict
		var Runes []rune = []rune(Word)
		for _, Rune := range Runes {
			if CurrentNode.nodeMap == nil {
				CurrentNode.nodeMap = make(map[rune]*MapNode)
			}
			_, exist := CurrentNode.nodeMap[Rune]
			if !exist {
				CurrentNode.nodeMap[Rune] = new(MapNode)
				CurrentNode.nodeMap[Rune].endOfWord = false
			}
			CurrentNode = CurrentNode.nodeMap[Rune]
		}
		CurrentNode.endOfWord = true
	}
	return
}

func (WordDict *WordDictionary) AddWordsFromFile(WordsFilePath string) (err error) {
	file, err := os.Open(WordsFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	var Words []string
	Words, err = readLines(file)
	if err != nil {
		return
	}

	err = WordDict.AddWords(Words)
	return

}

func readLines(file *os.File) (lines []string, err error) {
	var (
		part   []byte
		prefix bool
	)
	lines = make([]string, 0)
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))

	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}

	return
}
