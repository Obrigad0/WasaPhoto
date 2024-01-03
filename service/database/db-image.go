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
	err2 := db.c.QueryRow("SELECT iId,data FROM image WHERE author = ? ORDER BY data DESC LIMIT 1", i.Author).Scan(&id)
	if err2 != nil {
		// Errore nell'esecuzione della Query
		return -1, err2
	}

	return id, nil
}

// DeleteImage() elimina l'immagine caricata dall'utente
func (db *appdbimpl) DeleteImage(author User, iId Image) error {

	_, err := db.c.Exec("DELETE FROM image WHERE author = ? AND iId = ? ", author.UId, iId.IId)
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
	err := db.c.QueryRow("SELECT * FROM image WHERE iId = ? AND author NOT IN (SELECT banner FROM ban WHERE banned = ?)", iId.IId, requester.UId).Scan(&image)
	// visto il vincolo in tabella [ban] CHECK (banner != banned) l'utente potra' sempre vedere le proprie foto
	if err != nil {
		// Errore nell'esecuzione della Query
		return image, err
	}
	return image, err

}
