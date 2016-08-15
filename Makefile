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

build-test-container:
	@cd docker; docker build --rm -t go-rundeck-test:2.6.9 .; cd -

run-test-container:
	@docker run -d -p 127.0.0.1:4440:4440 --name go-rundeck-test -t go-rundeck-test:2.6.9

stop-test-container:
	@docker stop go-rundeck-test
	@docker rm go-rundeck-test

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
