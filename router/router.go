package router

import (
    "carver_golang_blog/api/controller"
    "github.com/gin-gonic/gin"
)

func CarverRouter() *gin.Engine {
    router := gin.Default()

    api := router.Group("api")
    {
        api.POST("login", controller.Login)//用户登录
    }

    return router
}
