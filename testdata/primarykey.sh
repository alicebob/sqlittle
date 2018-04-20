#!/bin/bash
set -eu


DB=primarykey.sqlite

rm -f $DB
(
    echo "CREATE TABLE words (word varchar NOT NULL PRIMARY KEY);"
    echo "BEGIN;"
    for w in $( cat words.txt ); do
        echo "INSERT INTO words VALUES (\"$w\");"
    done
    echo "COMMIT;"
) | sqlite3 --batch $DB
