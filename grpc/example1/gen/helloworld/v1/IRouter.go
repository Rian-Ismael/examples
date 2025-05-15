package v1

import "github.com/gin-gonic/gin"

type RouterGroup interface {
	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}
