package models

import (
	"bytes"
	"fmt"
	"gin_blogweb/config"
	"gin_blogweb/database"
	"github.com/jmoiron/sqlx"
	"html/template"
	"strconv"
	"strings"
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

func MakeHomeBlocks(article []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range (article) {
		homeParam := &HomeBlockParam{
			Id:         art.Id,
			Title:      art.Title,
			Tags:       art.Tags,
			Short:      art.Short,
			Content:    art.Content,
			Author:     art.Author,
			CreateTime: art.CreateTime,
			Link:       "/show" + strconv.Itoa(art.Id),
			UpdateLink: "/article/update?id=" + strconv.Itoa(art.Id),
			DeleteLink: "/article/delete?id=" + strconv.Itoa(art.Id),
			IsLogin:    isLogin,
		}

		t, _ := template.ParseFiles("views/home_block.html")
		buffer := bytes.Buffer{}
		_ = t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}

func CreateTagsLinks (tags string) (tagLink []TagLink){
	tagParams := strings.Split(tags, "&")
	for _, tag := range (tagParams){
		tagLink = append(tagLink, TagLink{
			TagName: "tag",
			TagUrl:  "/?tag=" + tag,
		})
	}
	return tagLink
}
