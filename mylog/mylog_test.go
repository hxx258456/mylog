// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.

package mylog

import (
	"bytes"
	"testing"
)

func TestLog(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out)
		log.Log().Msg("")
		if got, want := decodeIfBinaryToString(out.Bytes()), "{}\n"; got != want {
			t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
		}
	})

	t.Run("one-field", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out)
		log.Log().Str("foo", "bar").Msg("")
		if got, want := decodeIfBinaryToString(out.Bytes()), `{"foo":"bar"}`+"\n"; got != want {
			t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
		}
	})

	t.Run("two-field", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := New(out)
		log.Log().
			Str("foo", "bar").
			Int("n", 123).
			Msg("")
		if got, want := decodeIfBinaryToString(out.Bytes()), `{"foo":"bar","n":123}`+"\n"; got != want {
			t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
		}
	})
}
