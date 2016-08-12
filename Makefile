BINARIES := $(shell find src/ -maxdepth 1 -type d -name 'rundeck-*' -exec sh -c 'echo $(basename {})' \;)
BINLIST := $(subst src/,,$(BINARIES))

ifeq ($(TRAVIS_BUILD_DIR),)
	GOPATH := $(GOPATH)
else
	GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
endif

all: clean deps test rundeck $(BINLIST)

deps:
	@go get -t -d ./...

test: deps
	@go test rundeck.v12 -v
	@go test rundeck.v13 -v
	@go test rundeck.v17 -v

rundeck: deps
	@mkdir -p bin/
	@go install rundeck.v17

binaries: $(BINLIST)

$(BINLIST): deps
	@echo $@
	@go install $@

clean:
	@rm -rf bin/ pkg/ #src/github.com src/gopkg.in src/golang.org

.PHONY: all clean deps test rundeck $(BINLIST)
