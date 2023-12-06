package database


// lo stream dell'utente e' stato messo all'interno di db-user

//  FollowUser() permette all'utente user di seguire userToFollow
func (db *appdbimpl) FollowUser(user integer, userToFollow integer) error {		//A e' seguito da B
	_, err := db.c.Exec("INSERT INTO follow (A,B) VALUES (?, ?)", userToFollow, user)

	if err != nil{
		// Errore nell'esecuzione della Query
		return err
	}
	//nessun Errore 
	return nil
}
// UnfollowUser() permette all'utente di smettere di seguire followed
func (db *appdbimpl) UnfollowUser(user integer, followed integer) error {	//A e' seguito da B
	_, err := db.c.Exec("DELETE FROM follow WHERE A = ? AND B = ?", followed, user)

	if err != nil{
		// Errore nell'esecuzione della Query
		return err
	}
	//nessun Errore 
	return nil
}


//get my follower

func (db *appdbimpl) GetFollowingList(user integer) ([]integer,error) {		//A e' seguito da B

	rows, err := db.c.Query("SELECT A FROM follow WHERE B = ?  ", user)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()

	var following []integer	

	for rows.Next() {
		var user integer
		err := rows.Scan(&user)
		if err != nil {
			// Errore nella scansione della riga
			return nil, err
		}
		following = append(banList, user)
	}

	if rows.Err()!= nil {
		// Errore
		return nil, err
	}

	return following, nil

}


func (db *appdbimpl) GetFollowerList(user integer) ([]integer,error) {		//A e' seguito da B

	rows, err := db.c.Query("SELECT B FROM follow WHERE A = ?  ", user)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()

	var follower []integer	

	for rows.Next() {
		var user integer
		err := rows.Scan(&user)
		if err != nil {
			// Errore nella scansione della riga
			return nil, err
		}
		follower = append(banList, user)
	}

	if rows.Err()!= nil {
		// Errore
		return nil, err
	}

	return follower, nil

}