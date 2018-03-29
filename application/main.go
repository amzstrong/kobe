package main

import (
    "kobe/application/router"
    "kobe/lib"
)

func main() {
    lib.Start()
    r := router.InitRouter()
    //r.Run(":8989")
    lib.Run(r, ":8989")
}