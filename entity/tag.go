package entity

type Tag struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type TagClient struct {
	Name  string `json:"tagName"`
	Count int64  `json:"count"`
}
