#!/bin/bash
set -eu

rm -f magic.sqlite magic.tmp
sqlite3 --batch magic.tmp <<HERE
CREATE TABLE hello (who varchar(255));
INSERT INTO hello VALUES ("world");
INSERT INTO hello VALUES ("universe");
INSERT INTO hello VALUES ("town");
HERE

sed -e 's/SQLite/FooBar/' magic.tmp > magic.sqlite
rm -f magic.tmp
