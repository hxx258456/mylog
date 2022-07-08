// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.

package benchtest

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ScottMansfield/nanolog"
	"github.com/rs/zerolog"
)

func BenchmarkCompareToStdlib(b *testing.B) {
	b.Run("Nanolog", func(b *testing.B) {
		lw := nanolog.New()
		lw.SetWriter(ioutil.Discard)
		h := lw.AddLogger("foo thing bar thing %i64. Fubar %s foo. sadfasdf %u32 sdfasfasdfasdffds %u32.")
		args := []interface{}{int64(1), "string", uint32(2), uint32(3)}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lw.Log(h, args...)
		}
	})
	b.Run("Stdlib", func(b *testing.B) {
		args := []interface{}{int64(1), "string", uint32(2), uint32(3)}
		l := log.New(ioutil.Discard, "", 0)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			l.Printf("foo thing bar thing %d. Fubar %s foo. sadfasdf %d sdfasfasdfasdffds %d.", args...)
		}
	})
	b.Run("Zerolog_printf", func(b *testing.B) {
		args := []interface{}{int64(1), "string", uint32(2), uint32(3)}
		logger := zerolog.New(ioutil.Discard)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Log().Msgf("foo thing bar thing %d. Fubar %s foo. sadfasdf %d sdfasfasdfasdffds %d.", args...)

		}
	})
	b.Run("Zerolog_field", func(b *testing.B) {
		logger := zerolog.New(ioutil.Discard)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Log().Int64("foo thing bar thing", 1).
				Str("Fubar foo", "string").
				Uint32("sadfasdf", 2).
				Uint32("sdfasfasdfasdffds", 3).Msg("")
		}
	})
	b.Run("Zerolog_context", func(b *testing.B) {
		logger := zerolog.New(ioutil.Discard).With().
			Int64("foo thing bar thing", 1).
			Str("Fubar foo", "string").
			Uint32("sadfasdf", 2).
			Uint32("sdfasfasdfasdffds", 3).Logger()

		for i := 0; i < b.N; i++ {
			logger.Log().Msg("")
		}
	})
}
