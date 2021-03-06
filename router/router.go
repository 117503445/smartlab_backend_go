package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"smartlab/api"
	"smartlab/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(middleware.Cors())

	// 路由
	groupApi := r.Group("/api")
	{
		groupApi.POST("ping", api.Ping)
		bulletin := groupApi.Group("Bulletin")
		{
			bulletin.GET("", api.BulletinReadAll)
			bulletin.GET(":id", api.BulletinRead)
			auth := bulletin.Group("")
			auth.Use(middleware.JwtMiddleware.MiddlewareFunc())
			auth.Use(middleware.HasRole("admin"))
			{
				auth.POST("", api.BulletinCreate)
				auth.DELETE(":id", api.BulletinDelete)
				auth.PUT(":id",api.BulletinUpdate)
			}

		}
		dataLog := groupApi.Group("DataLog")
		{
			dataLog.POST("", api.DataLogCreate)
		}
		behaviorLog := groupApi.Group("BehaviorLog")
		{
			behaviorLog.POST("", api.BehaviorLogCreate)
			behaviorLog.GET("csv", api.BehaviorLogViewCSV)
		}
		feedback := groupApi.Group("feedback")
		{
			feedback.POST("", api.FeedbackCreate)
			feedback.GET("csv", api.FeedbackViewCSV)
		}
		wechat := groupApi.Group("wechat")
		{
			wechat.GET("/openid", api.GetOpenID)
		}

		user := groupApi.Group("user")
		{
			// 用户注册
			user.POST("", api.UserCreate)

			// 用户登录
			user.POST("login", middleware.JwtMiddleware.LoginHandler)

			auth := user.Group("")
			// 需要登录保护
			auth.Use(middleware.JwtMiddleware.MiddlewareFunc())
			{
				// User Routing
				auth.GET("me", api.UserRead)
				auth.PUT("", api.UserUpdate)
			}
		}
		admin := groupApi.Group("admin")
		// 需要登录保护
		admin.Use(middleware.JwtMiddleware.MiddlewareFunc())
		admin.Use(middleware.HasRole("admin"))
		{
			user := admin.Group("user")
			{
				user.GET(":id", api.AdminUserRead)
			}
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
