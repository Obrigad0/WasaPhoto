package database

//PostImage() posta l'immagine caricata dall'utente 
func (db *appdbimpl) PostImage(i Image)  error{
	// la struttura inviata ha nel campo iId nil, perche' non ancora creato
	_,err := db.c.Exec("INSERT INTO image (author,descrizione,file,data) VALUES (?,?,?,?)", i.author, i.descrizione, i.file, i.data)
	
	if err != nil {
		// Errore nell'esecuzione della Query
		return err
	}
	return nil
}

//DeleteImage() elimina l'immagine caricata dall'utente
func (db *appdbimpl) DeleteImage(author User, iId integer) error {

	_,err := db.c.Exec("DELETE FROM image WHERE author = ? AND iId = ? ",author, iId)
	if err != nil {
		// Errore nell'esecuzione della Query
		return err
	}
		return nil

}

// GetImage() recupera l'immagine richiesta dall'utente, se e' autorizzato
func (db *appdbimpl) GetImage(requester integer, iId integer) (Image,error) {
	//variabile dove andiamo ad inserire l'immagine richiesta
	var image Image
	// la query ritornera' l'immagine solamente se l'utente richiedente non e' stato bannato dal """postatore""" (author) dell'immagine
	err := db.c.QueryRow("SELECT * FROM image WHERE iId = ? AND author NOT IN (SELECT banner FROM ban WHERE banned = ?)",iId,requester).Scan(&image)
	// visto il vincolo in tabella [ban] CHECK (banner != banned) l'utente potra' sempre vedere le proprie foto
	if err != nil {
		// Errore nell'esecuzione della Query
		return err
	}
		return nil

}