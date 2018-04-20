#!/bin/bash
set -eu

# shuf /usr/share/dict/american-english | head -n 1000 > words.txt

DB=words.sqlite

rm -f $DB
(
    echo "CREATE TABLE words (word varchar, length int);"
    echo "BEGIN;"
    for w in $( cat words.txt ); do
        echo "INSERT INTO words VALUES (\"$w\", length(\"$w\"));"
    done
    echo "CREATE INDEX words_index_1 ON words (word);"
    echo "CREATE INDEX words_index_2 ON words (length, word);"
    echo "COMMIT;"
) | sqlite3 --batch $DB
