package manager

import "github.com/ggilmore/Go-Hangman/utilities"

type WordManager struct {
	wordsToAdd         chan string
	currentWordsAndIDs chan []string

	currentWordsList []string
	currentWordsSet  utilities.StringSet
}

func (wm *WordManager) Serve() {
	for {
		word := <-wm.wordsToAdd
		if !wm.currentWordsSet[word] {
			wm.currentWordsList = append(wm.currentWordsList, word)
			wm.currentWordsSet[word] = true
		}

		wordListCopy := make([]string, len(wm.currentWordsList))
		copy(wordListCopy, wm.currentWordsList)
		wm.currentWordsAndIDs <- wordListCopy
	}
}
