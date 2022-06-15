mylog
==========

# 介绍
一个简单的golang日志框架，支持服务器模式和本地模式，底层仍然使用golang的`log`包，在其基础上提供了以下功能:
- 在线修改具体函数的日志级别(仅在服务器模式下支持)
- 日志异步输出(仅在服务器模式下支持)
- 日志滚动，按天滚动，当天文件按配置的size滚动(仅在服务器模式下支持)
- 支持日志同时输出到文件与控制台
- 日志格式输出代码文件路径，代码行数，以及调用方函数包路径信息

本地模式与golang的`log`包基本相同，不具备日志级别在线修改、异步输出、文件滚动功能。

# 使用
mylog的使用很简单，直接依赖即可使用，默认使用本地模式，如果要使用服务器模式，只需要在代码中添加mylog的配置与初始化即可。

在对应的代码中使用:
```
import "github.com/hxx258456/mylog"
...

func test() {
    ...
    mylog.Debug("这是一条测试消息")
}
```

然后在相关工程的`go.mod`文件中添加:
```
github.com/hxx258456/mylog latest
```

# [License]

[License]: http://license.coscl.org.cn/MulanPSL2
Copyright (c) 2022 hxx258456
github.com/hxx258456/mylog is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2. 
You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2 
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.  
See the Mulan PSL v2 for more details.