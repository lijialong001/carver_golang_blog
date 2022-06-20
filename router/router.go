package router
import (
"github.com/gin-gonic/gin"
"carver_golang_blog/handler"
)
func CarverRouter() *gin.Engine {

    router := gin.Default()
    router.POST("/api/login", handler.Login)
    return router
}
