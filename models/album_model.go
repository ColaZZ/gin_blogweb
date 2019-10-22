package models

import "gin_blogweb/database"

type Album struct {
	Id         int    `db:"id""`
	FilePath   string `db:"filepath"`
	Filename   string `db:"filename"`
	Status     int    `db:"status"`
	CreateTime int64  `db:"createtime"`
}

func InsertAlbum(album Album) (int64, error) {
	return database.ModifyDB("insert into album(filepath, filename, status, createtime) values(?,?,?,?)",
		album.FilePath, album.Filename, album.Status, album.CreateTime)
}

func FindAllAlbums() (album []Album, err error) {
	sqlStr := "select id, filepath, filename, status, createtime"
	err = db.Select(&album, sqlStr)
	if err != nil {
		return nil, err
	}
	return album, err

}
