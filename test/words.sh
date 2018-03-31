#!/bin/bash
set -eu

# shuf /usr/share/dict/american-english | head -n 1000 > words.txt

rm -f words.sqlite
sqlite3 --batch words.sqlite <<HERE
CREATE TABLE words (word varchar, length int);
HERE
for w in $( cat words.txt ); do
    echo "INSERT INTO words VALUES (\"$w\", length(\"$w\"));"
done | sqlite3 --batch words.sqlite

sqlite3 --batch words.sqlite <<HERE
CREATE INDEX words_index_1 ON words (word);
CREATE INDEX words_index_2 ON words (length, word);
HERE
