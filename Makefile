BINARIES := $(shell find cmd/ -maxdepth 1 -type d -name 'rundeck*' -exec sh -c 'echo $(basename {})' \;)
BINLIST := $(subst src/,,$(BINARIES))

ifeq ($(TRAVIS_BUILD_DIR),)
	GOPATH := $(GOPATH)
else
	GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
endif

all: clean bindata test $(BINLIST)

test:
	@script/test

bindata:
	@cd ./pkg/rundeck/responses/testdata;  go-bindata -pkg testdata -o testdata.go *.json *.yaml *.txt *.aclpolicy; cd -

build-test-container:
	@cd docker; docker build --rm -t go-rundeck-test:2.10.1 .; cd -

run-test-container:
	@docker run -d -p 127.0.0.1:4440:4440 --name go-rundeck-test -t go-rundeck-test:2.10.1

stop-test-container:
	@docker stop go-rundeck-test
	@docker rm go-rundeck-test

binaries: $(BINLIST)

$(BINLIST):
	@echo $@
	@go install ./$@

clean:
	@rm -rf bin/

.PHONY: all clean bindata test $(BINLIST)
