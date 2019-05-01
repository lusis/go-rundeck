BINARIES := $(shell find cmd/ -maxdepth 1 -type d -name 'rundeck*' -exec sh -c 'echo $(basename {})' \;)
BINLIST := $(subst src/,,$(BINARIES))
RUNDECK_DEB_VERSION := 3.0.21.20190424-1.201904242241
#RUNDECK_DEB_VERSION := 3.0.9

all: clean bindata test $(BINLIST)

test:
	@script/test

bindata:
	@CGO_ENABLED=0 go run ./cmd/maketestdata/main.go

build-test-container:
	@cd docker; docker build --rm --build-arg RDECK_VER=$(RUNDECK_DEB_VERSION) -t go-rundeck-test:$(RUNDECK_DEB_VERSION) .; cd -

run-test-container:
	@docker run -d -p 4440:4440 --name go-rundeck-test -t go-rundeck-test:$(RUNDECK_DEB_VERSION)

stop-test-container:
	@docker stop go-rundeck-test
	@docker rm go-rundeck-test

binaries: $(BINLIST)

$(BINLIST):
	CGO_ENABLED=0 go install ./$@

clean:
	@rm -rf bin/

.PHONY: all clean bindata test $(BINLIST)
