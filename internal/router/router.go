package router

import (
	"order/internal/api"
	"order/internal/middleware"
	"order/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 初始化服务
	adminService := service.NewAdminService()
	demandService := service.NewDemandService()
	customerService := service.NewCustomerService()

	// 初始化处理器
	helloHandler := api.NewHelloHandler()
	adminHandler := api.NewAdminHandler(adminService)
	demandHandler := api.NewDemandHandler(demandService)
	customerHandler := api.NewCustomerHandler(customerService)

	r.GET("/", helloHandler.Hello)

	// API路由组
	v1 := r.Group("/api/v1")
	{
		// 公开接口
		v1.POST("/auth/login", adminHandler.Login)

		// 需要认证的接口
		auth := v1.Group("admin")
		auth.Use(middleware.JWTAuth())
		{
			auth.POST("", adminHandler.CreateAdmin)
			auth.GET("", adminHandler.GetAdminList)
			auth.GET("/:id/profile", adminHandler.GetAdminProfile)
			auth.PUT("/:id", adminHandler.UpdateAdmin)
			auth.DELETE("/:id", adminHandler.DeleteAdmin)
		}

		// 需求相关接口
		demand := v1.Group("/demand")
		demand.Use(middleware.JWTAuth())
		{
			demand.POST("", demandHandler.CreateDemand)
			demand.GET("", demandHandler.GetAllDemands)
			demand.GET("/:id", demandHandler.GetDemand)
			demand.PUT("/:id", demandHandler.UpdateDemand)
			demand.DELETE("/:id", demandHandler.DeleteDemand)
		}

		// 用户相关接口
		customer := v1.Group("/customer")
		customer.Use(middleware.JWTAuth())
		{
			customer.POST("", customerHandler.CreateCustomer)
			customer.GET("/:id", customerHandler.GetCustomer)
			customer.PUT("/:id", customerHandler.UpdateCustomer)
			customer.DELETE("/:id", customerHandler.DeleteCustomer)
		}
	}
}
