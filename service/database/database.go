/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	Name      string `json:"name"`
	UId       int    `json:"userId"`
	Followers []int  `json:"follower"`
	Following []int  `json:"following"`
	Ban       []int  `json:"banList"`
}

type Image struct {
	IId         int       `json:"imgId"`
	File        string    `json:"file"`
	Descrizione string    `json:"descrizione"`
	Like        []int     `json:"like"`
	Comments    []Comment `json:"comments"`
	Data        time.Time `json:"data"`
	Author      string    `json:"author"`
}

type Comment struct {
	CId       int    `json:"idComment"`
	Text      string `json:"text"`
	Commenter int    `json:"commenter"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// Tutte le informazioni sulle funzioni sono scritte nei file corrispondenti

	// FILE db-user.go
	ChangeUserName(user User, newName User) error
	GetUserName(user User) (string, error)
	GetUser(uId User) (User, error)
	CreateUser(nome User) (int, error)
	Access(nome User) (int, error)
	GetStream(uId User) ([]Image, []User, error)

	// FILE db-image.go
	PostImage(i Image) (int, error)
	DeleteImage(author User, iId Image) error
	GetImage(requester User, iId Image) (Image, error)

	// FILE db-follow.go
	FollowUser(user User, userToFollow User) error
	UnfollowUser(user User, followed User) error
	GetFollowingList(userr User) ([]int, error)
	GetFollowerList(userr User) ([]int, error)

	// FILE db-ban.go
	BanUser(banner User, banned User) error
	UnbanUser(banner User, banned User) error
	GetBanList(user User) ([]User, error)
	IsBanned(A User, B User) (bool, error)
	GetBanListVINT(user User) ([]int, error)

	// FILE db-like.go
	GetLikesList(image Image) ([]User, error)
	LikePhoto(user User, image Image) error
	UnlikePhoto(user User, image Image) error

	// FILE db-comments.go
	GetComments(image Image) ([]Comment, error)
	CommentPhoto(image Image, commento Comment) error
	UncommentPhoto(commento Comment) error
	GetTheCommenter(commento Comment) (int, error)
	// POSSIBILI FUNZIONI
	// GET FOLLOWER

	// Ping controlla se il database e' "Online" attivo, se non lo e' torna errore
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	// attiva le chiavi primarie
	_, errPramga := db.Exec(`PRAGMA foreign_keys= ON`)
	if errPramga != nil {
		return nil, fmt.Errorf("error setting pragmas: %w", errPramga)
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='user';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		err = createDB(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

// crea le tabelle del db
func createDB(db *sql.DB) error {
	// A SEGUITO DA B
	tables := [6]string{

		`CREATE TABLE IF NOT EXISTS user (
				uId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(16) UNIQUE NOT NULL
				CHECK (uId > 0)
			);`,

		`CREATE TABLE IF NOT EXISTS ban (
				banner INTEGER NOT NULL,
				banned INTEGER NOT NULL,
				PRIMARY KEY (banner,banned),
				FOREIGN KEY(banner) REFERENCES user (uId) ON DELETE CASCADE,
				FOREIGN KEY(banned) REFERENCES user (uId) ON DELETE CASCADE,
				CHECK (banner != banned)  -- Vincolo per evitare lo stesso utente
				);`,

		`CREATE TABLE IF NOT EXISTS follow (
				-- A e' seguito da B  
				A INTEGER NOT NULL,	
				B INTEGER NOT NULL,
				PRIMARY KEY (A,B),
				FOREIGN KEY(A) REFERENCES user (uId) ON DELETE CASCADE,
				FOREIGN KEY(B) REFERENCES user (uId) ON DELETE CASCADE,
				CHECK (A != B)  -- Vincolo per evitare lo stesso utente
				);`,

		`CREATE TABLE IF NOT EXISTS image (
				imgId INTEGER PRIMARY KEY AUTOINCREMENT,
				author INTEGER NOT NULL,
				descrizione VARCHAR(280),
				data DATETIME NOT NULL,
				FOREIGN KEY(author) REFERENCES user (uId) ON DELETE CASCADE
				);`,

		`CREATE TABLE IF NOT EXISTS  like (
				imgId INTEGER NOT NULL,
				uId INTEGER NOT NULL,
				PRIMARY KEY (imgId,uId),
				FOREIGN KEY(imgId) REFERENCES image (imgId) ON DELETE CASCADE,
				FOREIGN KEY(uId) REFERENCES user (uId) ON DELETE CASCADE
				);`,

		`CREATE TABLE IF NOT EXISTS comments (
				idComment INTEGER PRIMARY KEY AUTOINCREMENT,
				imgId INTEGER NOT NULL,
				uId INTEGER NOT NULL,
				text VARCHAR(80) NOT NULL,
				FOREIGN KEY(imgId) REFERENCES image (imgId) ON DELETE CASCADE,
				FOREIGN KEY(uId) REFERENCES user (uId) ON DELETE CASCADE
				);`,
	}

	for i := 0; i < len(tables); i++ {

		sqlStmt := tables[i]
		_, err := db.Exec(sqlStmt)

		if err != nil {
			return err
		}
	}
	return nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
