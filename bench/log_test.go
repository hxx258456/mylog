// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.
package bench

import (
	"bytes"
	"log"
	"os"
	"testing"

	mlog "github.com/hxx258456/mylog/log"
)

func BenchmarkLogStr(b *testing.B) {
	out := bytes.Buffer{}
	log.SetOutput(&out)
	mlog.Output(os.Stdout)
	b.Run("log", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			log.Println("test")
		}
	})

	b.Run("mylog", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mlog.Print("test")
		}
	})
}
