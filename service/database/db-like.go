package database


// GetLikesList() ritorna la lista degli utenti che hanno messo like ad una foto specifica
func (db *appdbimpl) GetLikesList(image integer) ([]User, error) {

	rows, err := db.c.Query("SELECT u.uId, u.name FROM user u, like l WHERE l.imgId = ? AND u.uId = l.uId ", image)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()
	// Array che conterra' tutti gli utenti che hanno meso like
	var likeList []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.uId, &user.name)
		if err != nil {
			// Errore nella scansione della riga
			return nil, err
		}
		likeList = append(likeList, user)
	}
	if rows.Err() != nil {
		// Errore 
		return nil, err
	}
	// Nessun errore, ritorno la lista degli utenti che hanno messo like
	return likeList, nil

}

//LikePhoto() mette il like dell'utente all'immagine
func (db *appdbimpl) LikePhoto(user integer,image integer) error {
	_, err := db.c.Exec("INSERT INTO like (iId,uId) VALUES (?, ?)", image, user)
	if err != nil {
		// Errore 
		return err
	}
	return nil
}

//UnlikePhoto()
func (db *appdbimpl) UnlikePhoto(user integer,image integer) error {
	_, err := db.c.Exec("DELETE FROM like WHERE uId = ? AND iId = ?", user,image)
	if err != nil {
		// Errore 
		return err
	}
	return nil
}