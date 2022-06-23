package router
import (
"github.com/gin-gonic/gin"
"carver_golang_blog/api/controller"
)
func CarverRouter() *gin.Engine {

    router := gin.Default()
    router.POST("/api/login", controller.Login)
    return router
}
