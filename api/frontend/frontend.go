package frontend

import (
	"errors"
	"oversoul/simple-blog/api/models"
)

func GetAllPosts() []models.Post {
	all_posts := []models.Post{}
	for _, post := range models.DB {
		all_posts = append(all_posts, post)
	}
	return all_posts
}

func ShowPost(slug string) (*models.Post, error) {
	if post, ok := models.DB[slug]; ok {
		return &post, nil
	}
	return nil, errors.New("post not found.")
}
