package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{1, "Title 1", "Article 1 body"},
	article{2, "Title 2", "Article 2 body"},
}

func getAllArticles() []article {
	return articleList
}
