#!/bin/bash
set -eu

DB="alter.sqlite"
rm -f $DB
(
    echo "CREATE TABLE words (word varchar);"
    echo "BEGIN;"
    for w in $( cat words.txt ); do
        echo "INSERT INTO words VALUES (\"$w\");"
    done
    echo "COMMIT;"
    echo "ALTER TABLE words add column something int default 42;"
) | sqlite3 --batch $DB
