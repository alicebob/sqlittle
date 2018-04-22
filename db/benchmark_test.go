package db

import (
	"math/rand"
	"testing"
)

func Benchmark_RandomRowid(b *testing.B) {
	db, err := OpenFile("../testdata/words.sqlite")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	r := rand.New(rand.NewSource(42))

	table, err := db.Table("words")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := int64(r.Intn(1000) + 1)
		row, err := table.Rowid(n)
		if err != nil {
			b.Fatal(err)
		}
		if len(row) != 2 {
			b.Fatal("invalid row")
		}
	}
}

func Benchmark_RandomIndex(b *testing.B) {
	db, err := OpenFile("../testdata/words.sqlite")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	words := wordList(b)
	rand.New(rand.NewSource(42)).Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

	index, err := db.Index("words_index_1")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, w := range words {
			index.ScanEq(
				Key{w},
				func(r Record) bool {
					if have, want := r[0].(string), w; have != want {
						b.Errorf("have %v, want %v", have, want)
					}
					return true
				},
			)
		}
	}
}
