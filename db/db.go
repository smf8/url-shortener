package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"

	"log"
	// blank import for sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/smf8/url-shortener/model"
)

var db *sql.DB

func Migrate(filename string, migrationPath string) error {
	var err error
	db, err = sql.Open("sqlite3", filename+".db")
	if err != nil {
		return err
	}
	fsrc, err := (&file.File{}).Open("file://" + migrationPath)
	if err != nil {
		return err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	m, err := migrate.NewWithInstance(
		"file",
		fsrc,
		"sqlite3",
		driver,
	)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}

//NewDB uses go-migrate migration to create tables in links.db
func NewDB(filename string) error {
	var err error
	db, err = sql.Open("sqlite3", filename+".db")
	if err != nil {
		return err
	}
	e := db.Ping()
	if e != nil {
		return err
	}
	return nil
}

//CreateDB creates a database file if it's created
// useless as we are using go-migrate
func CreateDB(filename string) {
	// open or create a sqlite database file with sqlite3 driver
	var err error
	db, err = sql.Open("sqlite3", filename+".db")
	if err != nil {
		log.Fatal(err)
	}
	err = createTable()
	if err != nil {
		log.Fatal("Error on creating table", err)
	}
}

//createTable creates sql table
func createTable() error {
	statement := `CREATE TABLE IF NOT EXISTS links (hash text unique not null, url text, usage integer DEFAULT 0)`
	_, err := db.Exec(statement)
	return err
}

//AddLink inserts a link into the database
func AddLink(link model.Link) {
	// preparing a db transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// prepare insert query
	statement, err := tx.Prepare("INSERT OR IGNORE INTO links(hash,url) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	// filling insert query with link data
	res, err := statement.Exec(link.Hash, link.Address)
	if err != nil {
		log.Fatal(res, err)
	}
	tx.Commit()
	// close the statement after transaction is finished
	defer statement.Close()
}

//Close closes the database connection
func Close() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//GetLink retrieves a link from database and returns it
func GetLink(hash string) model.Link {
	l := new(model.Link)
	statement, err := db.Prepare("SELECT hash,url,usage FROM links WHERE hash = ?")
	if err != nil {
		log.Fatal(err)
	}
	res := statement.QueryRow(hash)
	res.Scan(&l.Hash, &l.Address, &l.UsedTimes)
	defer statement.Close()
	return *l
}

//DeleteLink deletes a link from sqlite database
func DeleteLink(hash string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec("DELETE FROM links WHERE hash = ?", hash)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func IncrementUsage(hash string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec("UPDATE links SET usage = usage +1 WHERE hash = ?", hash)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
