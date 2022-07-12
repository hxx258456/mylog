// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.

package mylog

import "testing"

func TestHome(t *testing.T) {
	dir, err := Home()
	t.Log(dir)
	if err != nil {
		t.Error(err)
	}
}
