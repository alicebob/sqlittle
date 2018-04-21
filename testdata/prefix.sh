#!/bin/bash
set -eu

# Like words.sqlite, but with an extra column with the prefix of the words.
# Once expression columns in indexes are supported this can be simplified.

DB=prefix.sqlite

rm -f $DB
(
    echo "CREATE TABLE words (prefix varchar not null, word varchar not null primary key, length int not null);"
    echo "BEGIN;"
    for w in $( cat words.txt ); do
        echo "INSERT INTO words VALUES (substr(\"$w\", 0, 4), \"$w\", length(\"$w\"));"
    done
    echo "CREATE INDEX words_prefix ON words (prefix);"
    echo "CREATE INDEX words_prefix_desc ON words (prefix DESC);"
    echo "CREATE INDEX words_length ON words (length, word);"
    echo "COMMIT;"
) | sqlite3 --batch $DB
