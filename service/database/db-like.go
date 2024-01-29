package database

// GetLikesList() ritorna la lista degli utenti che hanno messo like ad una foto specifica
func (db *appdbimpl) GetLikesList(image Image) ([]int, error) {

	rows, err := db.c.Query("SELECT uId FROM like WHERE imgId = ?", image.IId)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()
	// Array che conterra' tutti gli utenti che hanno meso like
	var likeList []int
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UId)
		if err != nil {
			// Errore nella scansione della riga
			return nil, err
		}
		likeList = append(likeList, user.UId)
	}
	if rows.Err() != nil {
		// Errore
		return nil, err
	}
	// Nessun errore, ritorno la lista degli utenti che hanno messo like
	return likeList, nil

}

// LikePhoto() mette il like dell'utente all'immagine
func (db *appdbimpl) LikePhoto(user User, image Image) error {
	_, err := db.c.Exec("INSERT INTO like (imgId,uId) VALUES (?, ?)", image.IId, user.UId)
	if err != nil {
		// Errore
		return err
	}
	return nil
}

// UnlikePhoto()
func (db *appdbimpl) UnlikePhoto(user User, image Image) error {
	_, err := db.c.Exec("DELETE FROM like WHERE uId = ? AND imgId = ?", user.UId, image.IId)
	if err != nil {
		// Errore
		return err
	}
	return nil
}
