#!/bin/bash
set -eu

rm -f single.sqlite
sqlite3 --batch single.sqlite <<HERE
CREATE TABLE hello (who varchar(255));
INSERT INTO hello VALUES ("world");
INSERT INTO hello VALUES ("universe");
INSERT INTO hello VALUES ("town");
HERE
