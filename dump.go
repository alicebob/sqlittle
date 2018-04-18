// dump table and index structure
// usage: go run dump.go ./testdata/single.sqlite
// +build never

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alicebob/sqlittle"
)

func main() {
	flag.Parse()
	for _, f := range flag.Args() {
		db, err := sqlittle.OpenFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s", f, err)
			continue
		}

		info, err := db.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s", f, err)
			continue
		}
		fmt.Printf("%s:\n%s", f, info)
	}
}
