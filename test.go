package main

import (
    "fmt"
    "kobe/application/model"
    "kobe/lib"
)

func main() {
    user := new(model.UserModel)
    res, _ := user.M().Fields("name").First()
    fmt.Println(res)

    fmt.Println(lib.TotalDefine)
}
