package wordmap

import (
	"sort"
	"github.com/uvgroovy/findwords/powerset"
)

type WordsMap map[string][]string


func toString(set []interface{}) string {

	str := make([]rune, len(set))
	for i, v := range set {
		str[i] = v.(rune)
	}

	return string(str)
}

type runeSlice []rune

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }



func wordToKey(word string) string {
	var wordLetters runeSlice
	for _, c := range word {
		wordLetters = append(wordLetters, c)
	}
	sort.Sort(wordLetters)

	return string(wordLetters)
}



func (words *WordsMap) AddWord(word string) {
	key := wordToKey(word)
	(*words)[key] = append((*words)[key], word)
}



func removeDups(sets [][]interface{}) []string {
	// map to bool = poor man's set
	set := make(map[string]bool)
	for _, word := range sets {
		set[wordToKey(toString(word))] = true
	}
	// get the keys
	words := make([]string, 0, len(set))
    for k := range set {
        words = append(words, k)
    }
	return words
}


func (words *WordsMap) GetWords(letters string) <-chan string {
	
	c := make(chan string)
	go func() {
		lettersarr := make([]interface{}, 0)
		for _, chr := range letters {
			lettersarr = append(lettersarr, chr)
		}


		for _, lettersSubSet := range removeDups(powerset.CreatePowerSet(lettersarr)) {
			if len(lettersSubSet) < 2 {
				continue
			}

			if subWords, ok := (*words)[lettersSubSet]; ok {
				for _, word := range subWords {
					c <- word
				} 
			}
		}
		close(c)
	
	}()
	return c
}
