package database

// PostImage() posta l'immagine caricata dall'utente
func (db *appdbimpl) PostImage(i Image) (int, error) {
	// la struttura inviata ha nel campo iId nil, perche' non ancora creato
	_, err := db.c.Exec("INSERT INTO image (author,descrizione,data) VALUES (?,?,?)", i.Author, i.Descrizione, i.Data)

	if err != nil {
		// Errore nell'esecuzione della Query
		return -1, err
	}

	var id int
	err2 := db.c.QueryRow("SELECT imgId FROM image WHERE author = ? ORDER BY data DESC LIMIT 1", i.Author).Scan(&id)
	if err2 != nil {
		// Errore nell'esecuzione della Query
		return -2, err2
	}

	return id, nil
}

// DeleteImage() elimina l'immagine caricata dall'utente
func (db *appdbimpl) DeleteImage(author User, iId Image) error {

	_, err := db.c.Exec("DELETE FROM image WHERE author = ? AND imgId = ? ", author.UId, iId.IId)
	if err != nil {
		// Errore nell'esecuzione della Query
		return err
	}
	return nil

}

// GetImage() recupera l'immagine richiesta dall'utente, se e' autorizzato
func (db *appdbimpl) GetImage(requester User, iId Image) (Image, error) {
	// variabile dove andiamo ad inserire l'immagine richiesta
	var image Image
	// la query ritornera' l'immagine solamente se l'utente richiedente non e' stato bannato dal """postatore""" (author) dell'immagine
	err := db.c.QueryRow("SELECT * FROM image WHERE imgId = ? ", iId.IId).Scan(&image)
	// visto il vincolo in tabella [ban] CHECK (banner != banned) l'utente potra' sempre vedere le proprie foto
	if err != nil {
		// Errore nell'esecuzione della Query
		return image, err
	}
	return image, nil

}

// GetImage() recupera tutte le immagini
func (db *appdbimpl) GetAllImage(requester User) ([]Image, error) {

	var images []Image

	rows, err := db.c.Query(" SELECT * from image WHERE author = ? ORDER BY imgId DESC", requester.UId)

	if err != nil {
		// Errore nell'esecuzione della Query
		return nil, err
	}
	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()

	for rows.Next() {
		var image Image
		err := rows.Scan(&image.IId, &image.Author, &image.Descrizione, &image.Data)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return images, nil
}
