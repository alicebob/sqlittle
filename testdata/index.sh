#!/bin/bash
set -eu

rm -f index.sqlite
sqlite3 --batch index.sqlite <<HERE
CREATE TABLE hello (who varchar(255));
INSERT INTO hello VALUES ("world");
INSERT INTO hello VALUES ("universe");
INSERT INTO hello VALUES ("town");
CREATE INDEX hello_index ON hello (who);
HERE
