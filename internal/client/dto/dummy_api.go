package dto

import "time"

type DummyApiUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Picture   string `json:"picture"`
	Title     string `json:"title"`
}

type DummyApiPostResponse struct {
	ID          string               `json:"id"`
	Text        string               `json:"text"`
	Image       string               `json:"image"`
	Likes       int                  `json:"likes"`
	Tags        []string             `json:"tags"`
	PublishDate time.Time            `json:"publishDate"`
	Owner       DummyApiUserResponse `json:"owner"`
}
type DummyAPIPaginatedResponse struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Total int `json:"total"`
}

type DummyAPIUsersResponse struct {
	DummyAPIPaginatedResponse
	Data []DummyApiUserResponse `json:"data"`
}

type DummyAPIPostsResponse struct {
	DummyAPIPaginatedResponse
	Data []DummyApiPostResponse `json:"data"`
}
