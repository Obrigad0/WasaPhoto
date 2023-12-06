package database

//BanUser() banna l'utente banned
func (db *appdbimpl) BanUser(banner integer, banned integer) error {
	_, err := db.c.Exec("INSERT INTO ban (banner,banned) VALUES (?, ?)", banner, banned)
	if err != nil {
		// Errore nell'esecuzione della query
		return err
	}
	// Nessun errore
	return nil
}

//UnbanUser() sbanna l'utente banned
func (db *appdbimpl) UnbanUser(banner integer, banned integer) error {
	_, err := db.c.Exec("DELETE FROM ban WHERE banner = ? AND banned = ?", banner, banned)
	if err != nil {
		// Errore nell'esecuzione della query
		return err
	}
	// Nessun errore
	return nil
}

//GetBanList() ritorna la lista degli utenti bannati dalll'utente
func (db *appdbimpl) GetBanList(user integer) ([]User,error) {
	rows, err := db.c.Query("SELECT u.uId, u.name FROM user u, ban b WHERE u.uId = b.banned AND b.banner = ?  ", user)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()
	// Array che conterra' tutti gli utenti bannati
	var banList []User	// sono riempiti solo la variabile per il nome e per l'id!!!

	for rows.Next() {
		var user User
		err := rows.Scan(&user.uId, &user.name)
		if err != nil {
			// Errore nella scansione della riga
			return nil, err
		}
		banList = append(banList, user)
	}

	if rows.Err()!= nil {
		// Errore
		return nil, err
	}

	// Nessun errore, ritorno la lista degli utenti bannati
	return banList, nil
}

//IsBanned() controlla se A e' stato bannato da B
func (db *appdbimpl) IsBanned(A integer,B integer) (bool,error) {

	var result bool
	err := db.c.QueryRow(`SELECT EXISTS ( SELECT 1 FROM ban WHERE banner = ? AND banned = ?) AS isBanned;`, B,A).Scan(&result)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	return result,nil
}



//GetBanList() ritorna la lista degli utenti bannati dalll'utente
func (db *appdbimpl) GetBanListVINT(user integer) ([]integer,error) {

	rows, err := db.c.Query("SELECT banned FROM ban WHERE banner = ?", user)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()
	// Array che conterra' tutti gli utenti bannati
	var banList []int

	// Scorrimento delle righe query
	for rows.Next() {
		var user int
		err := rows.Scan(&user)
		if err != nil {
			// Errore nella scansione della riga
			return nil, err
		}
		banList = append(banList, user)
	}

	if rows.Err() != nil {
		// Errore scorrimento risposa querty
		return nil, err
	}

	// Nessun errore, ritorno la lista degli utenti bannati
	return banList, nil

}
