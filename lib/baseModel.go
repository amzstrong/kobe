package lib

import (
    "github.com/gohouse/gorose"
)

type BaseModel struct {
    database  string             //要连接的库名
    table     string             //表
    TableReal string
    flag      string             //分表标记
    Con       *gorose.Connection //数据库连接实例
}

//初始化分表标志位 例如user_02
func (m *BaseModel) InitFlag(flag string) *BaseModel {
    m.flag = flag
    return m
}
//mode初始化
func (m *BaseModel)InitConnect(args ...interface{}) {
    if len(args) == 1 {
        panic("At least two args (database ,table)")
    } else if len(args) == 2 {
        if database, ok := args[0].(string); ok {
            m.database = database
        } else {
            panic("only string allowed of  database name")
        }
        if table, ok := args[1].(string); ok {
            if m.flag != "" {
                m.TableReal = table + m.flag
            } else {
                m.TableReal = table
            }
        } else {
            panic("only string allowed of  table name")
        }
    } else {
        panic("2  params need in InitDb method ")
    }

    con, _ := Dao(m.database)
    m.Con = &con
}

