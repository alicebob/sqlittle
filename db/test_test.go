package db

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

// wordList gives the contents of words.txt
func wordList(t testing.TB) []string {
	f, err := os.Open("./../testdata/words.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var words []string
	b := bufio.NewReader(f)
	for {
		w, err := b.ReadString('\n')
		if err == io.EOF {
			return words
		}
		if err != nil {
			t.Fatal(err)
		}
		words = append(words, strings.TrimRight(w, "\n"))
	}
}
