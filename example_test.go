package sqlittle

import (
	"fmt"
)

func ExampleDatabase_TableScan() {
	db, err := OpenFile("test/single.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.TableScan(
		"hello",
		func(rowid int64, rec Record) bool {
			fmt.Printf("row %d: %s\n", rowid, rec[0].(string))
			return false // we want all the rows
		},
	); err != nil {
		panic(err)
	}
	// output:
	// row 1: world
	// row 2: universe
	// row 3: town
}

func ExampleDatabase_TableRowid() {
	db, err := OpenFile("test/single.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row, err := db.TableRowid("hello", 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("row: %s\n", row[0].(string))
	// output:
	// row: universe
}

func ExampleDatabase_IndexScan() {
	// This code will iterate over all words in alphabetical order.
	// The `words` table has: CREATE INDEX words_index_1 ON words (word)
	db, err := OpenFile("test/words.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	i := 0
	if err := db.IndexScan(
		"words_index_1",
		func(rowid int64, rec Record) bool {
			fmt.Printf("row %d: %s\n", rowid, rec[0].(string))
			i++
			return i >= 10
		},
	); err != nil {
		panic(err)
	}
	// output:
	// row 329: Adams
	// row 123: Ahmadinejad
	// row 870: Alabaman
	// row 685: Algonquin
	// row 619: Amy
	// row 700: Andersen
	// row 900: Annette's
	// row 423: Antipas's
	// row 891: Arizonan
	// row 945: Artaxerxes's
}

func ExampleDatabase_IndexScanMin() {
	// This will iterate over all words in alphabetical order, starting from
	// the first record >= the given record.
	// The `words` table has: CREATE INDEX words_index_1 ON words (word)
	db, err := OpenFile("test/words.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.IndexScanMin(
		"words_index_1",
		Record{"wombat"},
		func(rowid int64, rec Record) bool {
			word := rec[0].(string)
			if word >= "y" {
				return true
			}
			fmt.Printf("%s\n", word)
			return false
		},
	); err != nil {
		panic(err)
	}
	// output:
	// wombat
	// workbook
	// world's
	// worsens
	// wristwatch's
	// writhing
	// wusses
}
