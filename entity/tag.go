package entity

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TagClient struct {
	Name string `json:"tagName"`
}
