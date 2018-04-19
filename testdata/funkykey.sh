#!/bin/sh

# non-rowid table with complicated primary key and complicated indexes

DB=funkykey.sqlite

rm -rf $DB

(
    cat <<HERE
CREATE TABLE fuz (
    a,
    b,
    c,
    d,
    primary key(c, a),
    unique(b),
    unique(b, c),
    unique(a, c)
) WITHOUT ROWID;
INSERT INTO fuz VALUES ("angle", "billiards", "crotchety", "delta");
INSERT INTO fuz VALUES ("algebraic", "begotten", "colder", "destinies");
INSERT INTO fuz VALUES ("allegory", "beagle", "consequent", "duffers");
HERE
) | sqlite3 --batch $DB

