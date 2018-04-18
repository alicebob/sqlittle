#!/bin/bash
set -eu

rm -f four.sqlite
sqlite3 --batch four.sqlite <<HERE
CREATE TABLE aap (who varchar(255));
CREATE TABLE noot (who varchar(255));
CREATE TABLE mies (who varchar(255));
CREATE TABLE vuur (who varchar(255));
INSERT INTO aap VALUES ("world");
INSERT INTO aap VALUES ("universe");
INSERT INTO aap VALUES ("town");
HERE

