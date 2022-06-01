package models

import (
	"errors"
	"strings"
)

var DB = map[string]Post{
	"my-first-post": Post{
		"my-first-post",
		"my first post",
		"dumb content of my first post",
	},
}

type Post struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	Body string `json:"body"`
}

func slugify(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func NewPost(name, body string) (post Post, err error) {
	slug := slugify(name)
	if _, ok := DB[slug]; ok {
		err = errors.New("post already exists.")
		return
	}

	DB[slug] = Post{slug, name, body}
	post = DB[slug]
	return
}

func UpdatePost(slug, name, body string) (err error) {
	if _, ok := DB[slug]; !ok {
		err = errors.New("Post not found.")
		return
	}

	DB[slug] = Post{slug, name, body}
	return
}

func DeletePost(slug string) (err error) {
	if _, ok := DB[slug]; !ok {
		err = errors.New("Post not found.")
		return
	}

	delete(DB, slug)
	return
}
