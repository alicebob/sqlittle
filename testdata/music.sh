#!/bin/sh

DB=music.sqlite

rm -rf $DB

(
    cat <<HERE
CREATE TABLE artists (
    id integer primary key autoincrement not null,
    name
);
CREATE TABLE albums (
    id integer primary key autoincrement not null,
    artist integer not null,
    name
);
CREATE TABLE tracks (
    id integer primary key not null,
    album integer not null,
    name,
    length
) WITHOUT ROWID;
CREATE INDEX albums_name ON albums (name);
CREATE INDEX tracks_length ON tracks (length);
INSERT INTO artists VALUES (1, 'The Beatles');
INSERT INTO albums VALUES (1, 1, 'Rubber Soul');
INSERT INTO albums VALUES (2, 1, 'Abbey Road');
INSERT INTO tracks VALUES (1, 1, 'Drive My Car', 2*60+25);
INSERT INTO tracks VALUES (2, 1, 'Norwegian Wood', 2*60+1);
INSERT INTO tracks VALUES (3, 1, 'You Wont See Me', 3*60+18);
INSERT INTO tracks VALUES (4, 2, 'Come Together', 4*60+19);
INSERT INTO tracks VALUES (5, 2, 'Something', 3*60+2);
INSERT INTO tracks VALUES (6, 2, 'Maxwells Silver Hammer', 3*60+27);
HERE
) | sqlite3 --batch $DB

