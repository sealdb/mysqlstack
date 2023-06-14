/*
 * mysqlstack
 *
 * Copyright (c) XeLabs
 * Copyright (c) 2023-2030 NeoDB Author
 * GPL License
 *
 */

package sqldb

import (
	"testing"
)

func TestConstants(t *testing.T) {
	var i byte
	for i = 0; i < COM_RESET_CONNECTION+2; i++ {
		CommandString(i)
	}
}
