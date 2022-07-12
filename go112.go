//go:build go1.12

// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.//go:build go1.12

package mylog

// Since go 1.12, some auto generated init functions are hidden from
// runtime.Caller.
const contextCallerSkipFrameCount = 2
