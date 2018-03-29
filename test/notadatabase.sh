#!/bin/bash
set -eu

# not a sql file
cat >notadatabase.sqlite <<HERE
long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file long enough file but not a sqlite file
HERE
