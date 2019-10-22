package models

import (
	"bytes"
	"fmt"
	"gin_blogweb/config"
	"gin_blogweb/database"
	"html/template"
	"strconv"
	"strings"
)

var artcileRowsNum = 0

type HomeBlockParam struct {
	Id         int    //`db:"id"`
	Title      string //`db:"title"`
	Tags       string //`db:"tags"`
	Short      string //`db:"short"`
	Content    string //`db:"content"`
	Author     string //`db:"author"`
	CreateTime int64  //`db:"createtime"`

	Link       string
	UpdateLink string
	DeleteLink string
	IsLogin    bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	PreLink  string
	NextLink string
	ShowPage string
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

func CreateTagsLinks(tags string) (tagLink []TagLink) {
	tagParams := strings.Split(tags, "&")
	for _, tag := range (tagParams) {
		tagLink = append(tagLink, TagLink{
			TagName: "tag",
			TagUrl:  "/?tag=" + tag,
		})
	}
	return tagLink
}

func ConfigHomeFooterPageCoder(page int) (pagecode HomeFooterPageCode) {
	num := GetArticleRowsNum()
	allPageNum := (num-1)/config.NUM + 1

	pagecode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	if page <= 1 {
		pagecode.HasPre = false
	} else {
		pagecode.HasPre = true
	}

	if page >= allPageNum {
		pagecode.HasNext = false
	} else {
		pagecode.HasNext = true
	}

	pagecode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pagecode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return
}

func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

func QueryArticleRowNum() int {
	sql := "select count(id) from article"
	rows := database.QueryRowDB(sql)
	num := 0
	_ = rows.Scan(&num)
	return num
}

func SetArticleRowsNum() {
	artcileRowsNum = QueryArticleRowNum()
}
