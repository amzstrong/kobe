package lib

import (
    "path/filepath"
    "strings"
    "os"
    "io/ioutil"
)


//返回上一级目录
func Dir(filePath string) string {
    path, _ := filepath.Abs(filePath)
    index := strings.LastIndex(path, string(os.PathSeparator))
    if index < 0 {
        panic("上一级目录不存在")
    }
    return path[:index]
}

//从文件读取数据,返回字符串
func ReadFile(filePath string) string {
    f, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    fc, err := ioutil.ReadAll(f)
    if err != nil {
        panic(err)
    }
    return string(fc)
}

