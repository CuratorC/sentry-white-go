// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	controllers "sentry-white-go/app/http/controllers/api/v1"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/api/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	// v1.Use(middlewares.LimitIP("200-H"))
	{
		ocGroup := v1.Group("originals")
		{
			oc := new(controllers.OriginalsController)
			ocGroup.GET("", oc.Index)
			ocGroup.GET("/:id", oc.Show)
			ocGroup.POST("", oc.Store)
			ocGroup.PUT("/:id", oc.Update)
			ocGroup.DELETE("/:id", oc.Delete)
		}

		pcGroup := v1.Group("projects")
		{
			pc := new(controllers.ProjectsController)
			pcGroup.GET("", pc.Index)
			pcGroup.GET("/:id", pc.Show)
			pcGroup.POST("", pc.Store)
			pcGroup.PUT("/:id", pc.Update)
			pcGroup.DELETE("/:id", pc.Delete)
		}

		rcGroup := v1.Group("robots")
		{
			rc := new(controllers.RobotsController)
			rcGroup.GET("", rc.Index)
			rcGroup.GET("/:id", rc.Show)
			rcGroup.POST("", rc.Store)
			rcGroup.PUT("/:id", rc.Update)
			rcGroup.DELETE("/:id", rc.Delete)
		}

		rpcGroup := v1.Group("responsible_people")
		{
			rpc := new(controllers.ResponsiblePeopleController)
			rpcGroup.GET("", rpc.Index)
			rpcGroup.GET("/:id", rpc.Show)
			rpcGroup.POST("", rpc.Store)
			rpcGroup.PUT("/:id", rpc.Update)
			rpcGroup.DELETE("/:id", rpc.Delete)
		}
	}
}
