#!/bin/bash
set -eu

# long records
rm -f overflow.sqlite
line=$(seq  -s"" "longline" 1000)
sqlite3 --batch overflow.sqlite <<HERE
CREATE TABLE mytable (myline varchar);
INSERT INTO mytable VALUES ("$line");
HERE
