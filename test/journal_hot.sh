#!/bin/bash
# kill sqlite while the transaction is in progress
# we want a dirty -journal file
set -eu

TABLE=journal_hot.sqlite

rm -f $TABLE ${TABLE}-journal

sqlite3 --batch $TABLE <<HERE
PRAGMA journal_mode=DELETE;
CREATE TABLE words (word);
BEGIN;
INSERT INTO words VALUES ("aap");
INSERT INTO words VALUES ("noot");
INSERT INTO words VALUES ("mies");
COMMIT;
HERE


(
    echo "BEGIN;"
    for w in $( cat words.txt ); do
        echo "INSERT INTO words VALUES (\"$w\");"
    done
    sleep 5 
) | sqlite3 --batch $TABLE &
sleep 1
kill -9 %1
