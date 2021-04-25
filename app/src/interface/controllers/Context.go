package controllers

import "mime/multipart"

type Context interface {
	GetHeader(key string) string
	Query(key string) string
	DefaultQuery(key string, defaultValue string) string
	Param(key string) string
	PostForm(key string) string
	PostFormArray(key string) []string
	GetPostForm(key string) (string, bool)
	JSON(code int, obj interface{})
	FormFile(name string) (*multipart.FileHeader, error)
}
