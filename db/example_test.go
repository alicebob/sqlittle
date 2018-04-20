package db

import (
	"fmt"
)

func ExampleTable_Scan() {
	db, err := OpenFile("../testdata/single.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	table, err := db.Table("hello")
	if err != nil {
		panic(err)
	}
	if err := table.Scan(
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

func ExampleTable_Scan_nonrowid() {
	db, err := OpenFile("../testdata/withoutrowid.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	table, err := db.NonRowidTable("words")
	if err != nil {
		panic(err)
	}
	i := 10
	if err := table.Scan(
		func(rec Record) bool {
			fmt.Printf("row %s\n", rec[0].(string))
			i--
			return i <= 0
		},
	); err != nil {
		panic(err)
	}
	// output:
	// row Adams
	// row Ahmadinejad
	// row Alabaman
	// row Algonquin
	// row Amy
	// row Andersen
	// row Annette's
	// row Antipas's
	// row Arizonan
	// row Artaxerxes's
}

func ExampleTable_Rowid() {
	db, err := OpenFile("../testdata/single.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	table, err := db.Table("hello")
	if err != nil {
		panic(err)
	}
	row, err := table.Rowid(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("row: %s\n", row[0].(string))
	// output:
	// row: universe
}

func ExampleTable_Rowid_nonrowid() {
	db, err := OpenFile("../testdata/withoutrowid.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	table, err := db.NonRowidTable("words")
	if err != nil {
		panic(err)
	}
	if err := table.ScanEq(Record{"awesomely"},
		func(r Record) bool {
			fmt.Printf("row: %v\n", r)
			return false
		},
	); err != nil {
		panic(err)
	}
	// output:
	// row: [awesomely 9]
}

func ExampleIndex_Scan() {
	// This code will iterate over all words in alphabetical order.
	// The `words` table has: CREATE INDEX words_index_1 ON words (word)
	db, err := OpenFile("../testdata/words.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	index, err := db.Index("words_index_1")
	if err != nil {
		panic(err)
	}
	i := 0
	if err := index.Scan(
		func(rec Record) bool {
			fmt.Printf("row %d: %s\n", rec[1].(int64), rec[0].(string))
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

func ExampleIndex_ScanMin() {
	// This will iterate over all words in alphabetical order, starting from
	// the first record >= the given record.
	// The `words` table has: CREATE INDEX words_index_1 ON words (word)
	db, err := OpenFile("../testdata/words.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	index, err := db.Index("words_index_1")
	if err != nil {
		panic(err)
	}
	if err := index.ScanMin(
		Record{"wombat"},
		func(rec Record) bool {
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

func ExampleIndex_ScanEq() {
	// Find records mathing a key.
	db, err := OpenFile("../testdata/words.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	index, err := db.Index("words_index_1")
	if err != nil {
		panic(err)
	}
	if err := index.ScanEq(
		Record{"wombat"},
		func(rec Record) bool {
			fmt.Printf("%v\n", rec) // word, rowid
			return false
		},
	); err != nil {
		panic(err)
	}
	// output:
	// [wombat 159]
}

func ExampleDatabase_Schema() {
	db, err := OpenFile("../testdata/words.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	schema, err := db.Schema("words")
	if err != nil {
		panic(err)
	}
	fmt.Printf("columns:\n")
	for _, c := range schema.Columns {
		fmt.Printf(" - %q is a %s\n", c.Column, c.Type)
	}
	fmt.Printf("available indexes:\n")
	for _, ind := range schema.Indexes {
		fmt.Printf(" - %q (", ind.Index)
		for i, c := range ind.Columns {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Print(c.Column)
		}
		fmt.Print(")\n")
	}
	// output:
	// columns:
	//  - "word" is a varchar
	//  - "length" is a int
	// available indexes:
	//  - "words_index_1" (word)
	//  - "words_index_2" (length, word)
}

func ExampleTable_Def() {
	db, err := OpenFile("../testdata/words.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	index, err := db.Index("words_index_2")
	if err != nil {
		panic(err)
	}
	ind, err := index.Def()
	if err != nil {
		panic(err)
	}
	for _, c := range ind.IndexedColumns {
		fmt.Printf("indexed column: %q (sortorder %s)\n", c.Column, c.SortOrder)
	}
	// output:
	// indexed column: "length" (sortorder ASC)
	// indexed column: "word" (sortorder ASC)
}
