/*
Copyright 2017 Google Inc.
Copyright 2023-2030 NeoDB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sqlparser

import (
	"strings"
	"testing"
)

func TestShow1(t *testing.T) {
	validSQL := []struct {
		input  string
		output string
	}{
		{
			input:  "show table status",
			output: "show table status",
		},
		{
			input:  "show table status from sbtest",
			output: "show table status from sbtest",
		},
		{
			input:  "show table status from sbtest where Name='t'",
			output: "show table status from sbtest where Name = 't'",
		},
		{
			input:  "show table status from sbtest like 't'",
			output: "show table status from sbtest like 't'",
		},
		{
			input:  "show create table t1",
			output: "show create table t1",
		},
		{
			input:  "show tables",
			output: "show tables",
		},
		{
			input:  "show full tables",
			output: "show full tables",
		},
		{
			input:  "show full tables from t1",
			output: "show full tables from t1",
		},
		{
			input:  "show full tables from t1 like '%mysql%'",
			output: "show full tables from t1 like '%mysql%'",
		},
		{
			input:  "show full tables where Table_type != 'VIEW'",
			output: "show full tables where Table_type != 'VIEW'",
		},
		{
			input:  "show tables from t1",
			output: "show tables from t1",
		},
		{
			input:  "show tables from t1 like '%mysql%'",
			output: "show tables from t1 like '%mysql%'",
		},
		{
			input:  "show databases",
			output: "show databases",
		},
		{
			input:  "show create database sbtest",
			output: "show create database sbtest",
		},
		{
			input:  "show create schema sbtest",
			output: "show create database sbtest",
		},
		{
			input:  "show storage engines",
			output: "show engines",
		},
		{
			input:  "show engines",
			output: "show engines",
		},
		{
			input:  "show status",
			output: "show status",
		},
		{
			input:  "show versions",
			output: "show versions",
		},
		{
			input:  "show processlist",
			output: "show processlist",
		},
		{
			input:  "show queryz",
			output: "show queryz",
		},
		{
			input:  "show txnz",
			output: "show txnz",
		},
		{
			input:  "show warnings",
			output: "show warnings",
		},
		{
			input:  "show warnings limit 1",
			output: "show warnings limit 1",
		},
		{
			input:  "show variables",
			output: "show variables",
		},
		{
			input:  "show variables like 'wait_timeout'",
			output: "show variables like 'wait_timeout'",
		},
		{
			input:  "show global variables like 'wait_timeout'",
			output: "show global variables like 'wait_timeout'",
		},
		{
			input:  "show variables where Variable_name='wait_timeout'",
			output: "show variables where Variable_name = 'wait_timeout'",
		},
		{
			input:  "show binlog events",
			output: "show binlog events",
		},
		{
			input:  "show binlog events limit 10",
			output: "show binlog events limit 10",
		},
		{
			input:  "show binlog events from gtid '20171225083823'",
			output: "show binlog events from gtid '20171225083823'",
		},
		{
			input:  "show binlog events from gtid '20171225083823' limit 1",
			output: "show binlog events from gtid '20171225083823' limit 1",
		},
		{
			input:  "show index from t1",
			output: "show index from t1",
		},
		{
			input:  "show indexes from t1",
			output: "show index from t1",
		},
		{
			input:  "show keys from t1",
			output: "show index from t1",
		},
		{
			input:  "show index in t1 in sbtest",
			output: "show index from sbtest.t1",
		},
		{
			input:  "show index from t1 from sbtest where Key_name='PRIMARY'",
			output: "show index from sbtest.t1 where Key_name = 'PRIMARY'",
		},
		{
			input:  "show columns from t1",
			output: "show columns from t1",
		},
		{
			input:  "show columns from t1 from sbtest",
			output: "show columns from sbtest.t1",
		},
		{
			input:  "show full columns in tt.t1 in sbtest",
			output: "show full columns from sbtest.t1",
		},
		{
			input:  "show columns from t1 like '%'",
			output: "show columns from t1 like '%'",
		},
		{
			input:  "show columns from t1 where `Key` = 'PRI'",
			output: "show columns from t1 where `Key` = 'PRI'",
		},
		{
			input:  "show full columns from t1",
			output: "show full columns from t1",
		},
		{
			input:  "show full columns from t1 like '%'",
			output: "show full columns from t1 like '%'",
		},
		{
			input:  "show full columns from t1 where `Key` = 'PRI'",
			output: "show full columns from t1 where `Key` = 'PRI'",
		},
		{
			input:  "show fields from t1",
			output: "show columns from t1",
		},
		{
			input:  "show fields from t1 like '%'",
			output: "show columns from t1 like '%'",
		},
		{
			input:  "show fields from t1 where `Key` = 'PRI'",
			output: "show columns from t1 where `Key` = 'PRI'",
		},
		{
			input:  "show full fields from t1",
			output: "show full columns from t1",
		},
		{
			input:  "show full fields from t1 like '%'",
			output: "show full columns from t1 like '%'",
		},
		{
			input:  "show full fields from t1 where `Key` = 'PRI'",
			output: "show full columns from t1 where `Key` = 'PRI'",
		},
		{
			input:  "show collation",
			output: "show collation",
		},
		{
			input:  "show collation where Collation='binary'",
			output: "show collation where `collation` = 'binary'",
		},
		{
			input:  "show collation like 'binary'",
			output: "show collation like 'binary'",
		},
		{
			input:  "show charset",
			output: "show charset",
		},
		{
			input:  "show char set",
			output: "show charset",
		},
		{
			input:  "show character set",
			output: "show charset",
		},
	}

	for _, show := range validSQL {
		sql := strings.TrimSpace(show.input)
		tree, err := Parse(sql)
		if err != nil {
			t.Errorf("input: %s, err: %v", sql, err)
			continue
		}

		// Walk.
		Walk(func(node SQLNode) (bool, error) {
			return true, nil
		}, tree)

		got := String(tree.(*Show))
		if show.output != got {
			t.Errorf("want:\n%s\ngot:\n%s", show.output, got)
		}
	}
}
