/*
   Copyright (c) 2022 hxx258456
   github.com/hxx258456/mylog is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package benchtest

import (
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/hxx25846/mylog/mylog"
)

func TestMain(m *testing.M) {
	mylog.ClearDir("testdata")
	m.Run()
}

func BenchmarkLogServer(b *testing.B) {
	setupServerOnce.Do(setupServer)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mylog.Debugf("测试写入日志: %d", i+1)
	}
}

func BenchmarkLogLocal(b *testing.B) {
	setupLocalOnce.Do(setupLocal)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mylog.Debugf("测试写入日志: %d", i+1)
	}
}

func BenchmarkLogGolang(b *testing.B) {
	setupGolangrOnce.Do(setupGolang)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Printf("测试写入日志: %d", i+1)
	}
}

var setupServerOnce sync.Once

func setupServer() {
	mylog.ClearDir("testdata")
	logConfig := &mylog.Config{
		LogForbidStdout:  true,
		LogFileDir:       "testdata",
		LogMod:           mylog.LOG_MODE_SERVER,
		LogLevelGlobal:   mylog.LOG_LEVEL_DEBUG,
		LogChnOverPolicy: mylog.LOG_CHN_OVER_POLICY_BLOCK,
		LogFileMaxSizeM:  100,
		LogChannelCap:    4096000,
		LogLevelCtlPort:  "19300",
	}
	mylog.InitLogger(logConfig)
	time.Sleep(3 * time.Second)
	mylog.Debug("准备测试日志文件")
}

var setupLocalOnce sync.Once

func setupLocal() {
	mylog.ClearDir("testdata")
	logConfig := &mylog.Config{
		LogForbidStdout: true,
		LogFileDir:      "testdata",
		LogMod:          mylog.LOG_MODE_LOCAL,
		LogLevelGlobal:  mylog.LOG_LEVEL_DEBUG,
	}
	mylog.InitLogger(logConfig)
	time.Sleep(3 * time.Second)
	mylog.Debug("准备测试日志文件")
}

var setupGolangrOnce sync.Once

func setupGolang() {
	mylog.ClearDir("testdata")
	currentLogFile, err := os.OpenFile("testdata/test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(currentLogFile)
	// 日志前缀时间戳格式
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("准备测试日志文件")
}

func TestClearLogs(t *testing.T) {
	mylog.ClearDir("testdata")
}
