package entity

type SnippetFirestore struct {
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	TagIds  []string `json:"tagIds"`
	Content string   `json:"content"`
}

type SnippetClient struct {
	Name    string   `json:"name"`
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}
