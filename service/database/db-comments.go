package database

//GetComments() ritorna tutti i commnenti di un'immagine
func (db *appdbimpl) GetComments(image Image) ([]Comment, error) {
	rows, err := db.c.Query("SELECT uId, text, commenter FROM comments c WHERE imgId = ?", image.IId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var commentList []Comment

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.CId, &comment.Text, &comment.Commenter)
		if err != nil {
			// Errore
			return nil, err
		}
		commentList = append(commentList, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	// Nessun errore, ritorno la lista dei commenti
	return commentList, nil
}

//CommentPhoto() inserisce un commento in un'immagine
func (db *appdbimpl) CommentPhoto(image Image, commento Comment) error {
	// commento non contiene l'id perche' verra' creato dal dbms
	_, err := db.c.Exec("INSERT INTO comment (uId,imgI,text) VALUES (?, ?, ?)", commento.Commenter, image.IId, commento.Text)
	if err != nil {
		//Errore
		return err
	}

	return nil
}

//UncommentPhoto() elimina il commento da un immagine
func (db *appdbimpl) UncommentPhoto(commento Comment) error {

	_, err := db.c.Exec("DELETE FROM comment WHERE idComment = ?", commento.CId)
	if err != nil {
		//Errore
		return err
	}

	return nil
}

func (db *appdbimpl) GetTheCommenter(commento Comment) (int, error) {

	var uId int
	// Esecuzione Query per selezionare il nome utente
	err := db.c.QueryRow(`SELECT uId FROM comments WHERE idComment = ?`, commento.CId).Scan(&uId)
	if err != nil {
		// Errore nell'esecuzione della Query
		return -1, err
	}
	//nessun Errore ritorno la tupla idUtente, errore (nil)
	return uId, nil

}
