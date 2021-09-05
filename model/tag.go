package model

type TagResponse struct {
	Tags []TagRemote `json:"tags"`
}

type TagRemote struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}
