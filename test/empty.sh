#!/bin/bash
set -eu

rm -f empty.sqlite
sqlite3 --batch empty.sqlite <<HERE
CREATE TABLE foo (foo varchar(1));
HERE
