package models

import "blogweb_gin/database"

type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	Createtime int64
}

//-------插入图片---------------
func InsertAlbum(album Album) (int64, error) {
	return database.ModifyDB("insert into album(filepath,filename,status,createtime)values(?,?,?,?)",
		album.Filepath, album.Filename, album.Status, album.Createtime)
}

//--------查询图片----------
func FindAllAlbums() ([]Album, error) {
	rows, err := database.QueryDB("select id,filepath,filename,status,createtime from album")
	if err != nil {
		return nil, err
	}
	var albums []Album
	for rows.Next() {
		id := 0
		filepath := ""
		filename := ""
		status := 0
		var createtime int64
		createtime = 0
		rows.Scan(&id, &filepath, &filename, &status, &createtime)
		album := Album{id, filepath, filename, status, createtime}
		albums = append(albums, album)
	}
	return albums, nil
}
