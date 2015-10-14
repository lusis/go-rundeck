BINARIES := $(shell find src/ -maxdepth 1 -type d -name 'rundeck-*' -exec sh -c 'echo $(basename {})' \;)
BINLIST := $(subst src/,,$(BINARIES))

ifeq ($(TRAVIS_BUILD_DIR),)
	GOPATH := $(GOPATH)
else
	GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
endif

all: clean test rundeck $(BINLIST)

test:
	@go test rundeck.v13 -v

rundeck:
	@mkdir -p bin/
	@go get ./... 
	@go install rundeck.v13

$(BINLIST):
	@echo $@
	@go install $@

clean:
	@rm -rf bin/ pkg/

.PHONY: all clean test rundeck $(BINLIST)
