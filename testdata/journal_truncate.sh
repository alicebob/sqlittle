#!/bin/bash
set -eu

TABLE=journal_truncate.sqlite

rm -f $TABLE ${TABLE}-journal

sqlite3 --batch $TABLE <<HERE
PRAGMA journal_mode=TRUNCATE;
CREATE TABLE words (word);
BEGIN;
INSERT INTO words VALUES ("aap");
INSERT INTO words VALUES ("noot");
INSERT INTO words VALUES ("mies");
COMMIT;
HERE
