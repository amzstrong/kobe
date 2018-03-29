package lib

import (
    "os"
    "io/ioutil"
    _"github.com/go-sql-driver/mysql"
    "github.com/gohouse/gorose"
    "encoding/json"
)

func open(args ...interface{}) (gorose.Connection, error) {
    return gorose.Open(args[0], args[1])
}

func getDbConfig(appPath string) interface{} {
    m := make(map[string]map[string]string)
    f, err := os.Open(appPath)
    if err != nil {
        panic(err)
    }
    js, err := ioutil.ReadAll(f)
    if err != nil {
        panic(err)
    }
    e := json.Unmarshal(js, &m)
    if e != nil {
        panic(e)
    }
    config := make(map[string]interface{})
    config["Connections"] = m
    config["Default"] = ""
    config["SetMaxOpenConns"] = 0
    config["SetMaxIdleConns"] = 1
    return config
}

func Dao(database string) (gorose.Connection, error) {
    configPath := "/Users/alex/go/src/vendertest/config/dev/mysql.json"
    config := getDbConfig(configPath)
    return open(config, database)
}
