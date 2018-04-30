package sqlittle_test

import (
	"fmt"

	"github.com/alicebob/sqlittle"
)

// Basic SELECT
func ExampleDB_Select() {
	db, err := sqlittle.Open("./testdata/music.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Select(
		"tracks",
		func(r sqlittle.Row) {
			var (
				name   string
				length int
			)
			_ = r.Scan(&name, &length)
			fmt.Printf("%s: %d seconds\n", name, length)
		},
		"name",
		"length",
	)
	// output:
	// Drive My Car: 145 seconds
	// Norwegian Wood: 121 seconds
	// You Wont See Me: 198 seconds
	// Come Together: 259 seconds
	// Something: 182 seconds
	// Maxwells Silver Hammer: 207 seconds

}

// SELECT in index order
func ExampleDB_IndexedSelect() {
	db, err := sqlittle.Open("./testdata/music.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.IndexedSelect(
		"tracks",
		"tracks_length",
		func(r sqlittle.Row) {
			var (
				name   string
				length int
			)
			_ = r.Scan(&name, &length)
			fmt.Printf("%s: %d seconds\n", name, length)
		},
		"name",
		"length",
	)
	// output:
	// Norwegian Wood: 121 seconds
	// Drive My Car: 145 seconds
	// Something: 182 seconds
	// You Wont See Me: 198 seconds
	// Maxwells Silver Hammer: 207 seconds
	// Come Together: 259 seconds
}

func ExampleDB_IndexedSelectEq() {
	db, err := sqlittle.Open("./testdata/music.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.IndexedSelectEq(
		"tracks",
		"tracks_length",
		sqlittle.Key{198},
		func(r sqlittle.Row) {
			var (
				name   string
				length int
			)
			_ = r.Scan(&name, &length)
			fmt.Printf("%s: %d seconds\n", name, length)
		},
		"name",
		"length",
	)
	// output:
	// You Wont See Me: 198 seconds
}

// SELECT a primary key
func ExampleDB_PKSelect() {
	db, err := sqlittle.Open("./testdata/music.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.PKSelect(
		"tracks",
		sqlittle.Key{4},
		func(r sqlittle.Row) {
			name, _ := r.ScanString()
			fmt.Printf("%s\n", name)
		},
		"name",
	)
	// output:
	// Come Together
}
