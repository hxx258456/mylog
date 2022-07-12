// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.

package mylog_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/hxx258456/mylog"

	"gopkg.in/natefinch/lumberjack.v2"
)

var filePath = "../testdata/test.log"

func TestLogStrcut(t *testing.T) {
	data := struct {
		Name string
		Age  int64
	}{
		Name: "bob",
		Age:  32,
	}
	log := mylog.New(os.Stdout)
	log.Print(data)
}

func TestLogToFile(t *testing.T) {
	data := struct {
		Name string
		Age  int64
	}{
		Name: "bob",
		Age:  32,
	}
	filePath := filePath
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		t.Error(err)
	}
	// colorWrite := NewConsoleWriter()
	log := mylog.New(file)
	log.Info().Timestamp().Interface("data", data).Msg("test")
}

// rotate log
func TestRotateLog(t *testing.T) {
	log := mylog.New(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    1,    // 单个文件大小
		MaxBackups: 3,    // 旧日志文件的数量
		MaxAge:     28,   // 日志存活时长
		Compress:   true, // 压缩
		LocalTime:  true,
	})
	for i := 0; i < 100000; i++ {
		log.Info().Timestamp().Msg("test")
	}
}

// clean test
func TestClean(t *testing.T) {
	t.Cleanup(func() {
		dir, err := ioutil.ReadDir("../testdata")
		if err != nil {
			t.Error(err)
		}
		for _, d := range dir {
			os.RemoveAll(path.Join("../testdata", d.Name()))
		}

	})
}

func BenchmarkLog(b *testing.B) {

	str := faker.Paragraph()
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		b.Error(err)
	}
	// colorWrite := NewConsoleWriter()
	log := mylog.New(file).With().Timestamp().Caller().Logger()
	b.Run("stack", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			log.Info().Str("message", str).Send()
		}
	})

	b.Run("error", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			log.Error().Err(errors.New(str)).Send()
		}
	})
}
