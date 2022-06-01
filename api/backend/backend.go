package backend

import (
	"oversoul/simple-blog/api/models"
)

type CreatePostForm struct {
	Name string `json:"name" form:"name" binding:"required"`
	Body string `json:"body" form:"body" binding:"required"`
}

type UpdatePostForm struct {
	Name string `json:"name" form:"name" binding:"required"`
	Body string `json:"body" form:"body" binding:"required"`
}

func CreateNewPost(form CreatePostForm) (err error) {
	_, err = models.NewPost(form.Name, form.Body)
	return
}

func UpdatePost(slug string, form UpdatePostForm) (err error) {
	err = models.UpdatePost(slug, form.Name, form.Body)
	return
}

func DeletePost(slug string) (err error) {
	err = models.DeletePost(slug)
	return
}
