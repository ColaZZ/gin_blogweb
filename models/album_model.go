package models

import "gin_blogweb/database"

type Album struct {
	Id         int
	FilePath   string
	Filename   string
	status     int
	CreateTime int64
}

func InsertAlbum(album Album) (int64, error){
	return database.ModifyDB("insert into album(filepath, filename, status, createtime) values(?,?,?,?)",
		album.FilePath, album.Filename, album.status, album.CreateTime)
}
