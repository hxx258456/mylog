// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.

package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorFormatting(t *testing.T) {
	assert.Equal(
		t,
		"\x1b[32mfoo\x1b[0m",
		Green.Add("foo"),
		"Unexpected colored output.",
	)
}

func BenchmarkColorFormat(b *testing.B) {
	b.Run("ADD", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Red.Add("foo")
		}
	})
}
