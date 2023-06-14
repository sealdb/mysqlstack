/*
 * mysqlstack
 *
 * Copyright (c) XeLabs
 * Copyright (c) 2023-2030 NeoDB Author
 * GPL License
 *
 */

package main

import (
	"fmt"

	"github.com/sealdb/mysqlstack/driver"
	"github.com/sealdb/mysqlstack/xlog"
)

func main() {
	log := xlog.NewStdLog(xlog.Level(xlog.INFO))
	address := fmt.Sprintf(":4407")
	client, err := driver.NewConn("mock", "mock", address, "", "")
	if err != nil {
		log.Panic("client.new.connection.error:%+v", err)
	}
	defer client.Close()

	qr, err := client.FetchAll("SELECT * FROM MOCK", -1)
	if err != nil {
		log.Panic("client.query.error:%+v", err)
	}
	log.Info("results:[%+v]", qr.Rows)
}
