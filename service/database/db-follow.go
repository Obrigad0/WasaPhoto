package database

// lo stream dell'utente e' stato messo all'interno di db-user

//  FollowUser() permette all'utente user di seguire userToFollow
func (db *appdbimpl) FollowUser(user User, userToFollow User) error { //A e' seguito da B
	_, err := db.c.Exec("INSERT INTO follow (A,B) VALUES (?, ?)", userToFollow.uId, user.uId)

	if err != nil {
		// Errore nell'esecuzione della Query
		return err
	}
	//nessun Errore
	return nil
}

// UnfollowUser() permette all'utente di smettere di seguire followed
func (db *appdbimpl) UnfollowUser(user User, followed User) error { //A e' seguito da B
	_, err := db.c.Exec("DELETE FROM follow WHERE A = ? AND B = ?", followed.uId, user.uId)

	if err != nil {
		// Errore nell'esecuzione della Query
		return err
	}
	//nessun Errore
	return nil
}

//get my follower

func (db *appdbimpl) GetFollowingList(userr User) ([]int, error) { //A e' seguito da B

	rows, err := db.c.Query("SELECT A FROM follow WHERE B = ?  ", userr.uId)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()

	var following []int

	for rows.Next() {
		var user int
		err := rows.Scan(&user)
		if err != nil {
			// Errore nella scansione della riga
			return nil, err
		}
		following = append(following, user)
	}

	if rows.Err() != nil {
		// Errore
		return nil, err
	}

	return following, nil

}

func (db *appdbimpl) GetFollowerList(userr User) ([]int, error) { //A e' seguito da B

	rows, err := db.c.Query("SELECT B FROM follow WHERE A = ?  ", userr.uId)
	if err != nil {
		// Errore nell'esecuzione della query
		return nil, err
	}

	// uso defer per ritardare l'operazione rows.close() fino alla fine della funzione
	defer rows.Close()

	var follower []int

	for rows.Next() {
		var user int
		err := rows.Scan(&user)
		if err != nil {
			// Errore nella scansione della riga
			return nil, err
		}
		follower = append(follower, user)
	}

	if rows.Err() != nil {
		// Errore
		return nil, err
	}

	return follower, nil

}
