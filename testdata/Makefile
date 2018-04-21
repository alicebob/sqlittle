.PHONY: all clean

ALL:= \
    alter.sqlite \
    empty.sqlite \
    four.sqlite \
    funkykey.sqlite \
    index.sqlite \
    journal_hot.sqlite \
    journal_persist.sqlite \
    journal_truncate.sqlite \
    magic.sqlite \
    music.sqlite \
    notadatabase.sqlite \
    overflow.sqlite \
    prefix.sqlite \
    primarykey.sqlite \
    single.sqlite \
    truncated.sqlite \
    values.sqlite \
    wal.sqlite \
    wal_crashed.sqlite \
    words.sqlite \
    withoutrowid.sqlite \
    zerolength.sqlite \


all: ${ALL}

%.sqlite: %.sh
	bash $<

clean:
	rm -rf *.sqlite
