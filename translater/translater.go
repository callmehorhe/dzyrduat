package translater

import (
	"sort"
	"strings"
)

type Dicts struct {
	yryssagDict []string
	ironDict    []string
}

func NewDicts(yryssag, iron []string) *Dicts {
	return &Dicts{
		yryssagDict: yryssag,
		ironDict:    iron,
	}
}

func (dict *Dicts) Translate(word, lang string) string {
	word = strings.ToLower(word)
	switch lang {
	case "yryssag":
		ans := sort.SearchStrings(dict.yryssagDict, word)
		return dict.yryssagDict[ans]
	case "iron":
		word = strings.Replace(word, "ае", "æ", -1)
		ans := sort.SearchStrings(dict.ironDict, word)
		return dict.ironDict[ans]
	default:
		return dict.Translate(word, "iron")
	}
}
