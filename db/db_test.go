package db

import (
	"github.com/smf8/url-shortener/model"
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	CreateDB("testdb")
	links := make([]model.Link, 5)
	links[0] = model.NewLink("http://google.com")
	links[1] = model.NewLink("http://yahoo.com")
	links[2] = model.NewLink("http://bing.com")
	links[3] = model.NewLink("http://gmail.com")
	links[4] = model.NewLink("http://aut.ac.ir")
	for _, link := range links {
		AddLink(link)
	}
	retrievedLinks := make([]model.Link, 5)
	for i := range retrievedLinks {
		retrievedLinks[i] = GetLink(links[i].Hash)
		if retrievedLinks[i].Address != links[i].Address {
			t.Error("Error on link validation")
		}
	}
	DeleteLink(links[0].Hash)
	DeleteLink(links[1].Hash)
	DeleteLink(links[2].Hash)
	row, err := db.Query("SELECT COUNT(*) FROM links")
	if err != nil {
		t.Error(err)
	}
	var i int
	row.Next()
	row.Scan(&i)
	if i != 2 {
		t.Error("Error on link deletion")
	}
	row.Close()

	IncrementUsage(links[4].Hash)
	IncrementUsage(links[4].Hash)
	l := GetLink(links[4].Hash)
	if l.UsedTimes != 2 {
		t.Error("Error on incrementing usage")
	}
	defer os.Remove("./testdb.db")
	defer Close()
}
