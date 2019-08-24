package db

import (
	"database/sql"
	"log"
	"os"
	// blank import for sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/smf8/url-shortener/model"
)

var db *sql.DB

//CreateDB creates a database file if it's created
func CreateDB(filename string) {
	// remove database file if it exists
	os.Remove(filename + ".db")

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
	statement := `CREATE TABLE IF NOT EXISTS links (hash text unique not null, url text)`
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
	statement, err := tx.Prepare("INSERT INTO links(hash,url) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	// filling insert query with link data
	res, err := statement.Exec(link.Address, link.Hash)
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
