# mylog
---

## Simple Logging Example
```go
package main

import (
    "github.com/hxx258456/mylog/mylog"
    "github.com/hxx258456/mylog/log"
)

func main() {
    // UNIX Time is faster and smaller than most timestamps
    mylog.TimeFieldFormat = mylog.TimeFormatUnix

    log.Print("hello world")
}
// Output: {"time":1516134303,"level":"debug","message":"hello world"}
```

## Rotate Logging Example
```go
package main

import (
    "github.com/hxx258456/mylog/mylog"
    "github.com/hxx258456/mylog/log"
    "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
    // UNIX Time is faster and smaller than most timestamps
    mylog.TimeFieldFormat = mylog.TimeFormatUnix

    log.Output(&lumberjack.Logger{
		Filename:   "../testdata/test.log", // 文件路径
		MaxSize:    1,    // 单个文件大小
		MaxBackups: 3,    // 旧日志文件的数量
		MaxAge:     28,   // 日志存活时长
		Compress:   true, // 压缩
		LocalTime:  true,
	})
    log.Print("hello world")
}

// Output: {"time":1516134303,"level":"debug","message":"hello world"}
```

## Context Logging Example

```go
package main

import (
    "github.com/hxx258456/mylog/mylog"
    "github.com/hxx258456/mylog/log"
)

func main() {
    // UNIX Time is faster and smaller than most timestamps
    mylog.TimeFieldFormat = mylog.TimeFormatUnix

    log.Debug().
        Str("Scale", "833 cents").
        Float64("Interval", 833.09).
        Msg("Fibonacci is everywhere")
    
    log.Debug().
        Str("Name", "Tom").
        Send()
}

// Output: {"level":"debug","Scale":"833 cents","Interval":833.09,"time":1562212768,"message":"Fibonacci is everywhere"}
// Output: {"level":"debug","Name":"Tom","time":1562212768}
```

## Reference
本项目参考了以下开源项目,并基于其代码做了二次开发,相对应开源作者表示感谢!

- `github.com/rs/zerolog`
- `github.com/ScottMansfield/nanolog`