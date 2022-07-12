// Copyright (c) 2022 hxx258456
// github.com/hxx258456/mylog is licensed under Mulan PSL v2.

// mylog/utils.go 通用工具
//
package mylog

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

// 获取当前用户Home目录
func Home() (string, error) {
	// 优先使用当前系统用户的Home目录
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}
	// 判断操作系统是否windows
	if runtime.GOOS == "windows" {
		return homeWindows()
	}
	return homeUnix()
}

func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("无法获取当前用户Home目录(unix)")
	}
	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("无法获取当前用户Home目录(windows)")
	}
	return home, nil
}

func B2s(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// S2b converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func S2b(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}
