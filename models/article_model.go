package models

import "gin_blogweb/database"

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64
}

func AddArticle(article Article) (int64, error) {
	row, _ := insertArticle(article)
	return row, nil
}

func insertArticle(article Article) (int64, error) {
	sqlStr := "insert into article(title, tags, short, content, author, createtime) values(?, ?, ?, ?, ?, ?)"
	row, _ := database.ModifyDB(sqlStr, article.Title, article.Tags, article.Short, article.Content, article.Author,
		article.CreateTime)
	return row, nil
}
