package main

import (
    "fmt"
    "vendertest/application/model"
    "vendertest/lib"
)

func main() {
    user := new(model.UserModel)
    res, _ := user.M().Fields("name").First()
    fmt.Println(res)

    fmt.Println(lib.TotalDefine)
}
