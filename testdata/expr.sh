#!/bin/bash
set -eu
TABLE=expr.sqlite

rm -f ${TABLE}
sqlite3 --batch ${TABLE} <<HERE
CREATE TABLE expr (name varchar(255));
CREATE INDEX expr_name ON expr (substr(name, 0, 10));
CREATE INDEX expr_where ON expr (name) WHERE name > "foo";
INSERT INTO expr values ("aap"), ("foo"), ("qqq"), ("longestnameever");
HERE
