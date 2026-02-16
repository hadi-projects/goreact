package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hadi-projects/go-react-starter/internal/handler"
	"github.com/hadi-projects/go-react-starter/internal/middleware"
)

func (r *Router) setupPrivateRoutes(
	v1 *gin.RouterGroup,
	authHandler handler.AuthHandler,
	userHandler handler.UserHandler,
	permissionHandler handler.PermissionHandler,
) {
	auth := v1.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", userHandler.Register)
	}

	users := v1.Group("/users")
	users.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		// User can access their own profile
		users.GET("/me", middleware.RoleGuard("user", "admin"), userHandler.Me)

		// Admin only for CRUD
		users.GET("", middleware.RoleGuard("admin"), userHandler.GetAll)
		users.PUT("/:id", middleware.RoleGuard("admin"), userHandler.Update)
		users.DELETE("/:id", middleware.RoleGuard("admin"), userHandler.Delete)
	}

	permissions := v1.Group("/permissions")
	permissions.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	permissions.Use(middleware.RoleGuard("admin"))
	{
		permissions.POST("", permissionHandler.Create)
		permissions.GET("", permissionHandler.GetAll)
		permissions.PUT("/:id", permissionHandler.Update)
		permissions.DELETE("/:id", permissionHandler.Delete)
	}
}
