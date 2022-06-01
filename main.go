package main

import (
	"net/http"

	"oversoul/simple-blog/api/backend"
	"oversoul/simple-blog/api/frontend"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		posts := frontend.GetAllPosts()
		ctx.JSON(http.StatusOK, posts)
	})

	r.GET("/posts/:slug", func(ctx *gin.Context) {
		post, err := frontend.ShowPost(ctx.Params.ByName("slug"))
		if err == nil {
			ctx.JSON(http.StatusOK, post)
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	})

	r.POST("/posts", func(ctx *gin.Context) {
		var form backend.CreatePostForm
		if err := ctx.ShouldBindWith(&form, binding.FormMultipart); err != nil {
			ctx.JSON(422, gin.H{"message": err.Error()})
			return
		}

		if err := backend.CreateNewPost(form); err != nil {
			ctx.JSON(422, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Post created."})
	})

	r.PUT("/posts/:slug", func(ctx *gin.Context) {
		var form backend.UpdatePostForm
		if err := ctx.ShouldBindWith(&form, binding.FormMultipart); err != nil {
			ctx.JSON(422, gin.H{"message": err.Error()})
			return
		}
		if err := backend.UpdatePost(ctx.Params.ByName("slug"), form); err != nil {
			ctx.JSON(404, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Post updated."})
	})

	r.DELETE("/posts/:slug", func(ctx *gin.Context) {
		if err := backend.DeletePost(ctx.Params.ByName("slug")); err != nil {
			ctx.JSON(422, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Post deleted."})
	})

	return r
}

func main() {
	r := setupRoutes()
	r.Run(":8000")
}
