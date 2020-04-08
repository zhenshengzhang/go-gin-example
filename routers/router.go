package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zzs/go-gin-example/pkg/setting"
	v1 "github.com/zzs/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/article/:id", v1.GetArticle)
		//创建文章
		apiv1.POST("/article", v1.CreateArticle)
		//更新指定文章
		apiv1.PUT("/article/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}
