.PHONY: all clean

ALL:=empty.sqlite \
    four.sqlite \
    index.sqlite \
    magic.sqlite \
    notadatabase.sqlite \
    overflow.sqlite \
    single.sqlite \
    truncated.sqlite \
    values.sqlite \
    words.sqlite \
    zerolength.sqlite \


all: ${ALL}

%.sqlite: %.sh
	bash $<

clean:
	rm -rf *.sqlite
