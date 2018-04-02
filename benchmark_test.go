package sqlittle

import (
	"math/rand"
	"testing"
)

func Benchmark_RandomRowid(b *testing.B) {
	db, err := OpenFile("test/words.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := rand.New(rand.NewSource(42))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := int64(r.Intn(1000)+1)
		row, err := db.TableRowid("words", n)
		if err != nil {
            b.Fatal(err)
		}
		if len(row) != 2 {
			b.Fatal("invalid row")
		}
	}
}
