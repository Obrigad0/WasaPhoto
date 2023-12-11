package api

import (
	"time"

	"github.com/Obrigad0/WasaPhoto/service/database"
)

type User struct {
	name      string `json:"name"`      // nome dell'utente
	uId       int    `json:"userId"`    // id dell'utente
	followers []int  `json:"follower"`  //insieme degli id degli utenti che seguono l'utente
	following []int  `json:"following"` // insieme degli id degli utenti seguti dall'utente
	ban       []int  `json:"banList"`   // insieme degli id degli utenti bannati dall'utente
}

type Image struct {
	iId         int       `json:"imgId"`
	file        string    `json:"file"`
	descrizione string    `json:"descrizione"`
	like        []int     `json:"like"`
	comments    []Comment `json:"comments"`
	data        time.Time `json:"data"`
	author      string    `json:"author"`
}

type Comment struct {
	cId       int    `json:"idComment"`
	text      string `json:"text"`
	commenter int    `json:"commenter"`
}

func (i Image) ToDatabase() database.Image {
	return database.Image{
		iId:         i.iId,
		file:        i.file,
		descrizione: i.descrizione,
		like:        i.like,
		data:        i.data,
		author:      i.author,
	}
}

func (u User) ToDatabase() database.User {
	return database.User{
		name:      u.name,
		uId:       u.uId,
		followers: u.followers,
		following: u.following,
		ban:       u.ban,
	}
}

func (c Comment) ToDatabase() database.Comment {
	return database.Comment{
		cId:       c.cId,
		text:      c.text,
		commenter: c.commenter,
	}
}
