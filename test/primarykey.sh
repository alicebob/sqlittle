#!/bin/bash
set -eu


DB=primarykey.sqlite

rm -f $DB
sqlite3 --batch $DB <<HERE
CREATE TABLE words (word varchar, PRIMARY KEY(word));
HERE
