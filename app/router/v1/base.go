package v1

import (
	"github.com/gin-gonic/gin"
	"wrblog-api-go/app/api/v1/api_base"
)

func WrBaseApi(apiV1 *gin.RouterGroup) {
	//获取列表
	BaseApi := apiV1.Group("/base")
	{
		ArticleApi := BaseApi.Group("/baseArticle")
		{
			ArticleApi.GET("/myPage", api_base.ApiMyArticlePageList)
			ArticleApi.GET("/page", api_base.ApiArticlePageList)
			ArticleApi.GET("/:articleId", api_base.ApiArticleInfo)
			ArticleApi.POST("/add", api_base.ApiAddArticle)
			ArticleApi.PUT("/edit", api_base.ApiEditArticle)
			ArticleApi.DELETE("/removes", api_base.ApiRemoveArticle)
			ArticleApi.DELETE("/deletes", api_base.ApiDeleteArticle)
		}
		ArticleGroupApi := BaseApi.Group("/baseArticleGroup")
		{
			ArticleGroupApi.GET("/myPage", api_base.ApiMyArticleGroupPageList)
			ArticleGroupApi.GET("/page", api_base.ApiArticleGroupPageList)
			ArticleGroupApi.GET("/:groupId", api_base.ApiArticleGroupInfo)
			ArticleGroupApi.POST("/add", api_base.ApiAddArticleGroup)
			ArticleGroupApi.PUT("/edit", api_base.ApiEditArticleGroup)
			ArticleGroupApi.DELETE("/removes", api_base.ApiRemoveArticleGroup)
			ArticleGroupApi.DELETE("/deletes", api_base.ApiDeleteArticleGroup)
		}
		ArticleLabelApi := BaseApi.Group("/baseArticleLabel")
		{
			ArticleLabelApi.GET("/myPage", api_base.ApiMyArticleLabelPageList)
			ArticleLabelApi.GET("/page", api_base.ApiArticleLabelPageList)
			ArticleLabelApi.GET("/:labelId", api_base.ApiArticleLabelInfo)
			ArticleLabelApi.POST("/add", api_base.ApiAddArticleLabel)
			ArticleLabelApi.PUT("/edit", api_base.ApiEditArticleLabel)
			ArticleLabelApi.DELETE("/removes", api_base.ApiRemoveArticleLabel)
			ArticleLabelApi.DELETE("/deletes", api_base.ApiDeleteArticleLabel)
		}
		ArticleLibraryApi := BaseApi.Group("/baseArticleLibrary")
		{
			ArticleLibraryApi.GET("/myPage", api_base.ApiMyArticleLibraryPageList)
			ArticleLibraryApi.GET("/page", api_base.ApiArticleLibraryPageList)
			ArticleLibraryApi.GET("/:libraryId", api_base.ApiArticleLibraryInfo)
			ArticleLibraryApi.POST("/add", api_base.ApiAddArticleLibrary)
			ArticleLibraryApi.PUT("/edit", api_base.ApiEditArticleLibrary)
			ArticleLibraryApi.DELETE("/removes", api_base.ApiRemoveArticleLibrary)
			ArticleLibraryApi.DELETE("/deletes", api_base.ApiDeleteArticleLibrary)
		}
	}
}
