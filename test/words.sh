#!/bin/bash
set -eu

DICT=/usr/share/dict/american-english

rm -f words.sqlite
sqlite3 --batch words.sqlite <<HERE
CREATE TABLE words (word varchar, length int);
HERE
for w in $( shuf $DICT | head -n 1000 ); do
    echo "INSERT INTO words VALUES (\"$w\", length(\"$w\"));"
done | sqlite3 --batch words.sqlite

sqlite3 --batch words.sqlite <<HERE
CREATE INDEX words_index_1 ON words (word);
CREATE INDEX words_index_2 ON words (length, word);
HERE
