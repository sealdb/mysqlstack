/*
 * mysqlstack
 *
 * Copyright (c) 2023-2030 NeoDB Author
 * GPL License
 *
 */
 
package sqlparser

import (
	"strings"
	"testing"
)

func TestNeoDB(t *testing.T) {
	validSQL := []struct {
		input  string
		output string
	}{
		// name, address, user, password.
		{
			input:  "neodb attach ('attach1', '127.0.0.1:6000', 'root', '123456')",
			output: "neodb attach ('attach1', '127.0.0.1:6000', 'root', '123456')",
		},
		{
			input:  "neodb attachlist",
			output: "neodb attachlist",
		},
		{
			input:  "neodb detach('attach1')",
			output: "neodb detach ('attach1')",
		},
		{
			input:  "neodb reshard db.t db.tt",
			output: "neodb reshard db.t to db.tt",
		},
		{
			input:  "neodb reshard db.t to a.tt",
			output: "neodb reshard db.t to a.tt",
		},
		{
			input:  "neodb reshard db.t as b.tt",
			output: "neodb reshard db.t to b.tt",
		},
		{
			input:  "neodb cleanup",
			output: "neodb cleanup",
		},
		{
			input:  "neodb xa recover",
			output: "neodb xa recover",
		},
		{
			input:  "neodb xa rollback",
			output: "neodb xa rollback",
		},
		{
			input:  "neodb xa commit",
			output: "neodb xa commit",
		},
		{
			input:  "neodb rebalance",
			output: "neodb rebalance",
		},
	}

	for _, exp := range validSQL {
		sql := strings.TrimSpace(exp.input)
		tree, err := Parse(sql)
		if err != nil {
			t.Errorf("input: %s, err: %v", sql, err)
			continue
		}

		// Walk.
		Walk(func(node SQLNode) (bool, error) {
			return true, nil
		}, tree)

		got := String(tree.(*NeoDB))
		if exp.output != got {
			t.Errorf("want:\n%s\ngot:\n%s", exp.output, got)
		}
	}
}
