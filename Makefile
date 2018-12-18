BINARIES := $(shell find cmd/ -maxdepth 1 -type d -name 'rundeck*' -exec sh -c 'echo $(basename {})' \;)
BINLIST := $(subst src/,,$(BINARIES))
RUNDECK_DEB_VERSION := 3.0.9.20181127-1.201811291844

ifeq ($(TRAVIS_BUILD_DIR),)
	GOPATH := $(GOPATH)
else
	GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
endif

all: clean bindata test $(BINLIST)

test:
#ifneq ($(TRAVIS_BUILD_DIR),)
#	@script/travis-integration
#endif
	@script/test

bindata:
	@go get -u github.com/jteeuwen/go-bindata/...
	@cd ./pkg/rundeck/responses/testdata;  go-bindata -pkg testdata -o testdata.go *.json *.yaml *.txt *.aclpolicy; cd -

build-test-container:
	@cd docker; docker build --rm --build-arg RDECK_VER=$(RUNDECK_DEB_VERSION) -t go-rundeck-test:$(RUNDECK_DEB_VERSION) .; cd -

run-test-container:
	@docker run -d -p 4440:4440 --name go-rundeck-test -t go-rundeck-test:$(RUNDECK_DEB_VERSION)

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
