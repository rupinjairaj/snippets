package entity

type SnippetFirestore struct {
	Id      int64    `json:"id"`
	TagIds  []string `json:"tagIds"`
	Content string   `json:"content"`
}

type SnippetClient struct {
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}
