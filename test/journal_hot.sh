#!/bin/bash
# Kill sqlite while the transaction is in progress,
# we want the dirty -journal file.
# We need a cache spill to force the valid journal file, hence the cache_size
# pragma.

set -eu

DB=journal_hot.sqlite

rm -f $DB ${DB}-journal

sqlite3 --batch $DB <<HERE
PRAGMA journal_mode=DELETE;
CREATE TABLE words (word);
BEGIN;
INSERT INTO words VALUES ("aap");
INSERT INTO words VALUES ("noot");
INSERT INTO words VALUES ("mies");
COMMIT;
HERE


(
    echo "PRAGMA cache_size=5;";
    echo "BEGIN;"
    for w in $( cat words.txt ); do
        echo "INSERT INTO words VALUES (\"$w\");"
    done
    sleep 5 
) | sqlite3 --batch $DB &
sleep 1
kill -9 %1
