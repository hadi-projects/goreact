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
	roleHandler handler.RoleHandler,
	logHandler handler.LogHandler,
	cacheHandler handler.CacheHandler,
	generatorHandler handler.GeneratorHandler,
	abcHandler handler.AbcHandler,
		xyzHandler handler.xyzHandler,
		ninaHandler handler.ninaHandler,
		sdsdsdHandler handler.sdsdsdHandler,
		akusajaHandler handler.akusajaHandler,
		makanHandler handler.MakanHandler,
	// [GENERATOR_INSERT_HANDLER_PARAM]
) {
	// Module Generator
	generator := v1.Group("/generator")
	generator.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		generator.POST("", middleware.PermissionGuard("create-module"), generatorHandler.Generate)
	}

	abc := v1.Group("/abc")
	abc.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		abc.POST("", abcHandler.Create)
		abc.GET("", abcHandler.GetAll)
		abc.GET("/:id", abcHandler.GetByID)
		abc.PUT("/:id", abcHandler.Update)
		abc.DELETE("/:id", abcHandler.Delete)
	}
		xyz := v1.Group("/xyz")
	xyz.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		xyz.POST("", xyzHandler.Create)
		xyz.GET("", xyzHandler.GetAll)
		xyz.GET("/:id", xyzHandler.GetByID)
		xyz.PUT("/:id", xyzHandler.Update)
		xyz.DELETE("/:id", xyzHandler.Delete)
	}
		nina := v1.Group("/nina")
	nina.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		nina.POST("", ninaHandler.Create)
		nina.GET("", ninaHandler.GetAll)
		nina.GET("/:id", ninaHandler.GetByID)
		nina.PUT("/:id", ninaHandler.Update)
		nina.DELETE("/:id", ninaHandler.Delete)
	}
		sdsdsd := v1.Group("/sdsdsdsdd")
	sdsdsd.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		sdsdsd.POST("", sdsdsdHandler.Create)
		sdsdsd.GET("", sdsdsdHandler.GetAll)
		sdsdsd.GET("/:id", sdsdsdHandler.GetByID)
		sdsdsd.PUT("/:id", sdsdsdHandler.Update)
		sdsdsd.DELETE("/:id", sdsdsdHandler.Delete)
	}
		akusaja := v1.Group("/akusaja")
	akusaja.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		akusaja.POST("", akusajaHandler.Create)
		akusaja.GET("", akusajaHandler.GetAll)
		akusaja.GET("/:id", akusajaHandler.GetByID)
		akusaja.PUT("/:id", akusajaHandler.Update)
		akusaja.DELETE("/:id", akusajaHandler.Delete)
	}
		makan := v1.Group("/makan")
	makan.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		makan.POST("", makanHandler.Create)
		makan.GET("", makanHandler.GetAll)
		makan.GET("/:id", makanHandler.GetByID)
		makan.PUT("/:id", makanHandler.Update)
		makan.DELETE("/:id", makanHandler.Delete)
	}
	// [GENERATOR_INSERT_GROUP]
	auth := v1.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", userHandler.Register)
	}

	logs := v1.Group("/logs")
	logs.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		// Internal permission check is handled inside GetLogs
		logs.GET("", logHandler.GetLogs)
	}

	users := v1.Group("/users")
	users.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		// User can access their own profile
		users.GET("/me", middleware.PermissionGuard("get-profile"), userHandler.Me)

		// Admin only for CRUD
		users.POST("", middleware.PermissionGuard("create-user"), userHandler.Create)
		users.GET("", middleware.PermissionGuard("get-user"), userHandler.GetAll)
		users.PUT("/:id", middleware.PermissionGuard("edit-user"), userHandler.Update)
		users.DELETE("/:id", middleware.PermissionGuard("delete-user"), userHandler.Delete)
	}

	permissions := v1.Group("/permissions")
	permissions.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	permissions.Use(middleware.PermissionGuard("get-permission")) // Assuming admin role has this
	{
		permissions.POST("", middleware.PermissionGuard("create-permission"), permissionHandler.Create)
		permissions.GET("", middleware.PermissionGuard("get-permission"), permissionHandler.GetAll)
		permissions.PUT("/:id", middleware.PermissionGuard("edit-permission"), permissionHandler.Update)
		permissions.DELETE("/:id", middleware.PermissionGuard("delete-permission"), permissionHandler.Delete)
	}

	roles := v1.Group("/roles")
	roles.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	roles.Use(middleware.PermissionGuard("get-role"))
	{
		roles.POST("", middleware.PermissionGuard("create-role"), roleHandler.Create)
		roles.GET("", middleware.PermissionGuard("get-role"), roleHandler.GetAll)
		roles.GET("/:id", middleware.PermissionGuard("get-role"), roleHandler.GetByID)
		roles.PUT("/:id", middleware.PermissionGuard("edit-role"), roleHandler.Update)
		roles.DELETE("/:id", middleware.PermissionGuard("delete-role"), roleHandler.Delete)
	}

	// Cache management
	cache := v1.Group("/cache")
	cache.Use(middleware.AuthMiddleware(r.config.JWT.Secret))
	{
		cache.DELETE("/clear", middleware.PermissionGuard("manage-cache"), cacheHandler.ClearAll)
	}
}
