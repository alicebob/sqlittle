#!/bin/bash
set -eu

DB=./withoutrowid.sqlite

rm -f $DB
(
    echo "BEGIN;"
    echo "CREATE TABLE words (word varchar primary key, length int) WITHOUT ROWID;"
    for w in $( cat words.txt ); do
        echo "INSERT INTO words VALUES (\"$w\", length(\"$w\"));"
    done
    echo "CREATE INDEX words_l ON words (length, word);"
    echo "COMMIT;"
) | sqlite3 --batch $DB
