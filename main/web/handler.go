package web

import "back-end-remote-server/main/db"

// @hello
// yo yo yo yo
type Pet struct {
	ID       int `json:"id"`
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	Status string `json:"status"`
}

type Pet2 struct {
	ID int `json:"id"`
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}

type APIStatus struct {
	Status string
}

type ServerResponse struct {
	servers []db.Server
	numbers int64
}

type WebsiteResponse struct {
	websites []db.Website
}
