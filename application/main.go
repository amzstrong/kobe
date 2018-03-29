package main

import (
    "vendertest/application/router"
    "vendertest/lib"
)

func main() {
    lib.Start()
    r := router.InitRouter()
    //r.Run(":8989")
    lib.Run(r, ":8989")
}