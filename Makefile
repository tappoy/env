PACKAGE=github.com/tappoy/env
WORKING_DIRS=tmp bin

SRC=*.go
BIN=tmp/$(shell basename $(CURDIR))
TEST=tmp/cover
DOC=Document.txt
COVER0=tmp/cover0

.PHONY: all clean fmt cover test lint

all: $(WORKING_DIRS) fmt $(BIN) test $(DOC) lint

clean:
	rm -rf $(WORKING_DIRS)

$(WORKING_DIRS):
	mkdir -p $(WORKING_DIRS)

fmt: $(SRC)
	go fmt ./...

go.sum: go.mod
	go mod tidy

$(BIN): $(SRC) go.sum
	go build -o $(BIN)

test: $(BIN)
	go test -v -tags=mock -vet=all -cover -coverprofile=$(COVER)

$(DOC): $(SRC)
	go doc -all . > $(DOC)

cover: $(TEST)
	grep "0$$" $(COVER) | sed 's!$(PACKAGE)!.!' | tee $(COVER0)

lint: $(BIN)
	go vet
