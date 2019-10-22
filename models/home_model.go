package models

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

