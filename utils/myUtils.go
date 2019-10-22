package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
)

// md5加密处理
func MD5(pwd string) string {
	return  fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
}

func SwitchMarkdownToHtml(content string) template.HTML{
	markdown := blackfriday.MarkdownCommon([]byte(content))
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		//fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		//fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}


