SOURCES=$(shell find -name "*.go" ! -name "*_test.go" ! -path "./cmd/*")

CMDS=bin/wt bin/wi

all: ${CMDS}

bin/%: cmd/%.go ${SOURCES}
	go build -o $@ $<

info:
	echo ${SOURCES}

clean:
	rm -rf bin/
