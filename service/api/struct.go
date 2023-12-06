package api

import{
	"time"
}

type User struct{
	name string `json:"name"` // nome dell'utente
	uId integer `json:"userId"`	// id dell'utente
	followers []integer `json:"follower"` //insieme degli id degli utenti che seguono l'utente
	following []integer `json:"following"` // insieme degli id degli utenti seguti dall'utente
	ban []integer `json:"banList"` // insieme degli id degli utenti bannati dall'utente
}

type Image struct{
	iId integer `json:"imgId"`
	file string `json:"file"`
	descrizione string `json:"descrizione"`
	like []integer `json:"like"`
	comments []Comment `json:"comments"`
	data  time.Time  `json:"data"`  
	author string	`json:"author"`  

}

type Comment struct{
	cId integer `json:"idComment"`
	text string `json:"text"`
	commenter integer `json:"commenter"`
}