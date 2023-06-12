export PATH := $(GOPATH)/bin:$(PATH)

fmt:
	go fmt ./...
	go vet ./...

test:
	go get github.com/stretchr/testify/assert
	@echo "--> Testing..."
	@$(MAKE) testxlog
	@$(MAKE) testsqlparser
	@$(MAKE) testsqldb
	@$(MAKE) testproto
	@$(MAKE) testpacket
	@$(MAKE) testdriver

testxlog:
	go test -v ./xlog
testsqlparser:
	go test -v ./sqlparser/...
testsqldb:
	go test -v ./sqldb
testproto:
	go test -v ./proto
testpacket:
	go test -v ./packet
testdriver:
	go test -v ./driver

goyacc:
	go build -v -o bin/goyacc tools/goyacc/main.go
	bin/goyacc -o sqlparser/sql.go sqlparser/sql.y

COVPKGS = ./sqlparser/... ./sqldb ./proto ./packet ./driver
coverage:
	# @$(MAKE) goyacc
	# go get github.com/pierrre/gotestcover
	go build -v -o bin/gotestcover tools/gotestcover/gotestcover.go
	bin/gotestcover -coverprofile=coverage.out -v $(COVPKGS)
	# TODO: If go version is bigger than 1.19, it will generate sqlparpser/yaccpar
	# in the coverage.out file.
	# To solve this problem completely, the sql.go must be regenerated with the new
	# version of goyacc, and change the way it is called in parser.go file.
	sed -i '/yaccpar/d' coverage.out
	go tool cover -html=coverage.out

.PHONY: fmt testcommon testproto testpacket testdriver coverage
