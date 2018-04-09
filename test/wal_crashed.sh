#!/bin/bash
set -eu

DB=wal_crashed.sqlite

rm -f $DB*

(
    cat <<HERE
PRAGMA journal_mode=WAL;
CREATE TABLE words (word varchar);
BEGIN;
HERE
    for w in $( cat words.txt ); do
        echo "INSERT INTO words VALUES (\"$w\");"
    done
    echo "COMMIT;"
    sleep 10
) | sqlite3 --batch $DB &
sleep 1
kill %1
