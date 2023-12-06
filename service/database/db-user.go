package database

// ChangeUserName() cambia il nome utente
func (db *appdbimpl) ChangeUserName(user integer, newName string) error {
	// Esecuzione Query per la modifica dell'user name
	_, err := db.c.Exec("UPDATE user SET name = ? WHERE uId = ?", newName, user)
	if err != nil {
		//Errore nella modifica del nome utente
		return err
	}
	// Nessun errore, torno nil
	return nil
}

// GetUserName() ritorna il nome utente
func (db *appdbimpl) GetUserName(user integer) (string, error) {
	// Variabile che conterra' il risultato della query (il nome utente)
	var name string
	// Esecuzione Query per selezionare il nome utente
	err := db.c.QueryRow(`SELECT name FROM user WHERE uId = ?`, user).Scan(&name)
	if err != nil{
		// Errore nell'esecuzione della Query
		return "",err
	}
	//nessun Errore ritorno la tupla nome utente, errore (nil)
	return name,nil
}

// GetUser() ritorna i dati dell'utente
func (db *appdbimpl) GetUser(uId integer) (User, error) {

	// Variabile che conterra' il risultato della query (l'utente)
	var utente User

	// Esecuzione Query per selezionare l'utente
	err := db.c.QueryRow(`
		SELECT u.uId, u.name
		FROM user u
		WHERE u.uId = ?
	`, uId).Scan(&utente.uId, &utente.name)

	if err != nil{
		// Errore nell'esecuzione della Query
		return nil,err
	}
	//nessun Errore ritorno la tupla utente, errore (nil)
	return utente,nil

}

// CreateUser() crea l'utente al primo accesso se l'account non esiste
func (db *appdbimpl) CreateUser(nome string)  (integer,error) {

	_, err := db.c.Exec("INSERT INTO user (name) VALUES (?)", nome)
	if err != nil {
		// Errore nell'esecuzione della Query
		return -1,err
	}
	//nessun Errore
	//Prelevo l'id dell'utente creato
	var uId integer
	err := db.c.QueryRow("SELECT uId FROM user WHERE name = ?", nome).Scan(&uId)
	if err != nil {
		// Errore nell'esecuzione della Query
		return -1,err
	}
	return uId,nil
}

// Access() quando accede l'utente inserisce il nome utente, la funzione ritorna il suo id
func (db *appdbimpl) Access(nome string) (integer, error) {
	// E' inizializzato -1 in modo tale che se la query non restituisce un risulato vuol dire che 
	// l'utente non esiste e va creato
	uId := -1
	err := db.c.QueryRow("SELECT uId FROM user WHERE name = ?", nome).Scan(&uId)
	if err != nil{
		// Errore nell'esecuzione della Query
		return -1 ,err
	}
	//nessun Errore ritorno la tupla utente, errore (nil)
	return uId,nil
}


//  GetStream() restituisce un array di immagini contenenti l'ultima foto postata da ogni utente e l'utente che l'ha postata,
func (db *appdbimpl) GetStream(uId integer) ([]Image,[]User error) {
	var image []Image
	var users []User

	rows, err := db.c.Query("
		SELECT i.imgId, i.author, i.descrizione, i.date, u.uId , u.name
		FROM image i JOIN user u ON i.author = u.uId
		WHERE i.author IN (
 							SELECT A
 							FROM follow
 							WHERE B = ?
							)
		ORDER BY i.date DESC
		LIMIT 1;
	", uId)

	//forse si puo fare piu efficiente
	if err != nil {
		// Errore nell'esecuzione della Query
		return nil, err
	}
	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()

	// Scorrimento delle righe restituite dalla query
	for rows.Next() {
		var img Image
		var user User
		err := rows.Scan(&img.imgId, &img.author, &img.descrizione, &img.date, &user.uId, &user.name)
		if err != nil {
			return nil, nil, err
		}
		images = append(images, img)
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, nil, err
	}

	// Nessun errore, ritorno lo stream
	return image, user, nil
}

//IdsTOUsers()	trasforma un array di integer che rappresentano id di utenti in un array di utenti
/*
func (db *appdbimpl) IdToUsers(id []integer) ([]User error) {

	var users []User

	for i:=0; i<len(id); i++{

		var utente User
		err := db.c.QueryRow("SELECT * FROM user WHERE uId = ?", id[i]).Scan(&utente)
		users = append(users, utente)

	}
}
*/

func (db *appdbimpl) UserExist(uId integer) (bool, error) {


}
