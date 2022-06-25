package main

import "carver_golang_blog/router"

func main() {
    //初始化go-gin，启动webserver
    r := router.CarverRouter()
    _ = r.Run(":8012")

}
