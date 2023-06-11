[![Build Status](https://travis-ci.org/sealdb/mysqlstack.png)](https://travis-ci.org/sealdb/mysqlstack) [![Go Report Card](https://goreportcard.com/badge/github.com/sealdb/mysqlstack)](https://goreportcard.com/report/github.com/sealdb/mysqlstack) [![codecov.io](https://codecov.io/gh/sealdb/mysqlstack/graphs/badge.svg)](https://codecov.io/gh/sealdb/mysqlstack/branch/main)

# mysqlstack

**_mysqlstack_** is an MySQL protocol library implementing in Go (golang).

Protocol is based on [mysqlproto-go](https://github.com/pubnative/mysqlproto-go) and [go-sql-driver](https://github.com/go-sql-driver/mysql)

## Running Tests

```
$ mkdir src
$ export GOPATH=`pwd`
$ go get -u github.com/sealdb/mysqlstack/driver
$ cd src/github.com/sealdb/mysqlstack/
$ make test
```

## Examples

1. **_examples/mysqld.go_** mocks a MySQL server by running:

```
$ go run example/mysqld.go
  2018/01/26 16:02:02.304376 mysqld.go:52:     [INFO]    mysqld.server.start.address[:4407]
```

2. **_examples/client.go_** mocks a client and query from the mock MySQL server:

```
$ go run example/client.go
  2018/01/26 16:06:10.779340 client.go:32:    [INFO]    results:[[[10 nice name]]]
```

## Status

mysqlstack is production ready.

## License

mysqlstack is released under the BSD-3-Clause License. See [LICENSE](https://github.com/sealdb/mysqlstack/blob/main/LICENSE)

