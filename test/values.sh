#!/bin/bash
set -eu

rm -f values.sqlite
sqlite3 --batch values.sqlite <<HERE
CREATE TABLE things (c varchar(255), i int, f float);
INSERT INTO things VALUES (NULL, 0, 0);
INSERT INTO things VALUES ("", 1, 0);
INSERT INTO things VALUES ("", 0, 0);
INSERT INTO things VALUES ("", 80, 0);
INSERT INTO things VALUES ("", -80, 0);
INSERT INTO things VALUES ("", 1<<14, 0);
INSERT INTO things VALUES ("", -1<<14, 0);
INSERT INTO things VALUES ("", 1<<20, 0);
INSERT INTO things VALUES ("", -1<<20, 0);
INSERT INTO things VALUES ("", 1<<30, 0);
INSERT INTO things VALUES ("", -1<<30, 0);
INSERT INTO things VALUES ("", 1<<42, 0);
INSERT INTO things VALUES ("", -1<<42, 0);
INSERT INTO things VALUES ("", 1<<53, 0);
INSERT INTO things VALUES ("", -1<<53, 0);
INSERT INTO things VALUES ("", 0, 3.14);
INSERT INTO things VALUES ("", 0, -3.14);
HERE
