#!/bin/bash
set -eu

rm -f long.sqlite
sqlite3 --batch long.sqlite <<HERE
CREATE TABLE bottles (wall varchar);
HERE
for i in {1..1000}; do
    echo "INSERT INTO bottles VALUES (\"bottles of beer on the wall $i\");"
done | sqlite3 --batch long.sqlite
