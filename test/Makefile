.PHONY: all clean

ALL:=empty.sqlite \
    four.sqlite \
    index.sqlite \
    journal_hot.sqlite \
    journal_persist.sqlite \
    journal_truncate.sqlite \
    magic.sqlite \
    notadatabase.sqlite \
    overflow.sqlite \
    single.sqlite \
    truncated.sqlite \
    values.sqlite \
    wal.sqlite \
    wal_crashed.sqlite \
    words.sqlite \
    zerolength.sqlite \


all: ${ALL}

%.sqlite: %.sh
	bash $<

clean:
	rm -rf *.sqlite
