package api

import (
	"time"

	"github.com/Obrigad0/WasaPhoto/service/database"
)

type User struct {
	Name      string `json:"name"`      // nome dell'utente
	UId       int    `json:"userId"`    // id dell'utente
	Followers []int  `json:"follower"`  // insieme degli id degli utenti che seguono l'utente
	Following []int  `json:"following"` // insieme degli id degli utenti seguti dall'utente
	Ban       []int  `json:"banList"`   // insieme degli id degli utenti bannati dall'utente
}

type Image struct {
	IId         int       `json:"imgId"`
	File        string    `json:"file"`
	Descrizione string    `json:"descrizione"`
	Like        []int     `json:"like"`
	Comments    []Comment `json:"comments"`
	Data        time.Time `json:"data"`
	Author      string    `json:"author"`
}

type Comment struct {
	CId       int    `json:"idComment"`
	Text      string `json:"text"`
	Commenter int    `json:"commenter"`
}

func (i Image) ToDatabase() database.Image {
	return database.Image{
		IId:         i.IId,
		File:        i.File,
		Descrizione: i.Descrizione,
		Like:        i.Like,
		Data:        i.Data,
		Author:      i.Author,
	}
}

func (u User) ToDatabase() database.User {
	return database.User{
		Name:      u.Name,
		UId:       u.UId,
		Followers: u.Followers,
		Following: u.Following,
		Ban:       u.Ban,
	}
}

func (c Comment) ToDatabase() database.Comment {
	return database.Comment{
		CId:       c.CId,
		Text:      c.Text,
		Commenter: c.Commenter,
	}
}
