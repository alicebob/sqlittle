#!/bin/bash
set -eu

rm -f empty.sql
sqlite3 --batch empty.sql <<HERE
CREATE TABLE foo (foo varchar(1));
DROP TABLE foo;
HERE

rm -f single.sql
sqlite3 --batch single.sql <<HERE
CREATE TABLE hello (who varchar(255));
INSERT INTO hello VALUES ("world");
INSERT INTO hello VALUES ("universe");
INSERT INTO hello VALUES ("town");
HERE

rm -f four.sql
sqlite3 --batch four.sql <<HERE
CREATE TABLE aap (who varchar(255));
CREATE TABLE noot (who varchar(255));
CREATE TABLE mies (who varchar(255));
CREATE TABLE vuur (who varchar(255));
INSERT INTO aap VALUES ("world");
INSERT INTO aap VALUES ("universe");
INSERT INTO aap VALUES ("town");
HERE

# truncated
head -c 50 single.sql > truncated.sql

# mess up the magic number
sed -e 's/SQLite/FooBar/' single.sql > magic.sql

# not a sql file
cat >notadatabase.sql <<HERE
long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file
HERE

# 0 bytes
echo -n >zerolength.sql
