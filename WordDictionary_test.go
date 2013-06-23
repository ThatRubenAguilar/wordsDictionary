// WordDictionary_test.go
package words

import (
	"testing"
)

func Test_Lookup(Test *testing.T) {
	var wordDict *WordDictionary = new(WordDictionary)
	str := []string{"word", "two", "new"}
	err := wordDict.AddWords(str)
	if err != nil {
		Test.Errorf("Error in Test_Lookup: Building Word Dictionary failed with error %s", err)
	}
	answers := []WordInfo{
		WordInfo{"", true, false},
		WordInfo{"tw", true, false},
		WordInfo{"Tw", true, false},
		WordInfo{"two", true, true},
		WordInfo{"twod", false, false}}
	LookupInDict(wordDict, answers, Test)
}

func Test_Copy(Test *testing.T) {
	var wordDict *WordDictionary = new(WordDictionary)
	str := []string{"word", "two", "new"}
	err := wordDict.AddWords(str)
	if err != nil {
		Test.Errorf("Error in Test_Copy: Building Word Dictionary failed with error %s", err)
	}
	answers := []WordInfo{
		WordInfo{"twange", true, true},
		WordInfo{"", true, false},
		WordInfo{"tw", true, false},
		WordInfo{"Tw", true, false},
		WordInfo{"two", true, true},
		WordInfo{"twod", false, false}}
	copyDict := wordDict.Copy()
	wordDict.AddWords([]string{"twange"})
	LookupInDict(wordDict, answers, Test)
	answers[0] = WordInfo{"twange", false, false}
	LookupInDict(copyDict, answers, Test)
}

func LookupInDict(wordDict *WordDictionary, answers []WordInfo, Test *testing.T) {
	for _, answer := range answers {
		LookupWord(Test, wordDict, answer.Word, answer)
	}
}

func LookupWord(Test *testing.T, WordDict *WordDictionary, Word string, answer WordInfo) {
	info, err := WordDict.Lookup(Word)
	if err != nil {
		Test.Errorf("Error in LookupWord: Word %s Lookup failed with error %s", Word, err)
	}
	if info != answer {
		Test.Errorf("Fail in LookupWord: Word %v Returned %v which is not the answer %v", Word, info, answer)
	}
}
