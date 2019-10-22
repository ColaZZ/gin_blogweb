package models

import (
	"fmt"
	"gin_blogweb/config"
	"gin_blogweb/database"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

var db *sqlx.DB

type Article struct {
	Id         int    `db:"id"`
	Title      string `db:"title"`
	Tags       string `db:"tags"`
	Short      string `db:"short"`
	Content    string `db:"content"`
	Author     string `db:"author"`
	CreateTime int64  `db:"createtime"`
}

func AddArticle(article Article) (int64, error) {
	row, _ := insertArticle(article)
	SetArticleRowsNum()
	return row, nil
}

func insertArticle(article Article) (int64, error) {
	sqlStr := "insert into article(title, tags, short, content, author, createtime) values(?, ?, ?, ?, ?, ?)"
	row, _ := database.ModifyDB(sqlStr, article.Title, article.Tags, article.Short, article.Content, article.Author,
		article.CreateTime)
	return row, nil
}

func FindArticleWithPage(page int) ([]Article, error) {
	return QueryArticleWithPage(page, config.NUM)
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticleWithCon(sql)
}

func QueryArticleWithCon(sql string) (articleList []Article, err error) {
	sql = "select id, title, tags, short, content, author, createtime " + sql
	err = db.Select(&articleList, sql)
	if err != nil {
		return articleList, err
	}
	return articleList, nil
}

func QueryArticleWithId(id int) (article Article) {
	sql := fmt.Sprintf("select id, title, tags, short, content, author, createtime where id = %d",
		strconv.Itoa(id))
	err := db.Select(&article, sql)
	if err != nil {
		return
	}
	return
}

func UpdateArticle(article Article) (int64, error) {
	return database.ModifyDB("update article set title=?,tags=?,short=?,content=?,author=?, createtime= ? "+
		"where id=?", article.Title, article.Tags, article.Short, article.Content, article.Author,
		article.CreateTime, article.Id)
}

func DeleteArticle(id int) (int64, error) {
	row, err := deleteArticleWithId(id)
	SetArticleRowsNum()
	return row, err
}

func deleteArticleWithId(id int) (int64, error) {
	return database.ModifyDB("delete from article where id =?", id)
}

func QueryArticleWithParam(param string) []string {
	rows, err := database.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}

	var paramList []string
	for rows.Next() {
		arg := ""
		_ = rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := "where tags like '%&" + tag + "%&'"
	sql += "or tags like '%&" + tag + "'"
	sql += "or tags like '" + tag + "%&'"
	sql += "or tags like '" + tag + "'"
	return QueryArticleWithCon(sql)
}
