#!/bin/bash
set -eu

rm -f truncated.sqlite truncated.tmp
sqlite3 --batch truncated.tmp <<HERE
CREATE TABLE hello (who varchar(255));
INSERT INTO hello VALUES ("world");
INSERT INTO hello VALUES ("universe");
INSERT INTO hello VALUES ("town");
HERE
head -c 50 truncated.tmp > truncated.sqlite
rm -f truncated.tmp
