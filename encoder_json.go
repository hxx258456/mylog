//go:build !binary_log

// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.

package mylog

// encoder_json.go file contains bindings to generate
// JSON encoded byte stream.

import (
	"github.com/hxx25846/mylog/internal/pkg/json"
)

var (
	_ encoder = (*json.Encoder)(nil)

	enc = json.Encoder{}
)

func init() {
	// using closure to reflect the changes at runtime.
	json.JSONMarshalFunc = func(v interface{}) ([]byte, error) {
		return InterfaceMarshalFunc(v)
	}
}

func appendJSON(dst []byte, j []byte) []byte {
	return append(dst, j...)
}

func decodeIfBinaryToBytes(in []byte) []byte {
	return in
}
