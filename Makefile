BINARIES = rundeck-get-history rundeck-get-job rundeck-list-jobs rundeck-list-executions rundeck-get-tokens rundeck-list-projects

GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
all: clean test rundeck rundeck-bin

test:
	@go test rundeck.v12 -v

rundeck:
	@mkdir -p bin/
	@go get ./... 
	@go install rundeck.v12

rundeck-bin:
	@mkdir -p bin/
	$(foreach bin,$(BINARIES),go install $(bin);)

clean:
	@rm -rf bin/ pkg/

.PHONY: all clean test rundeck rundeck-bin 
