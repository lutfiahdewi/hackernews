package links

import (
	"log"

	database "github.com/lutfiahdewi/hackernews/internal/pkg/db/mysql"
	"github.com/lutfiahdewi/hackernews/internal/users"
)

// #1: definition of struct that represent a link.
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// #2: function that insert a Link object into database and returns it’s ID.
func (link Link) Save() int64 {
	//#3:  our sql query to insert link into Links table. you see we used prepare here before db.
	//Exec, the prepared statements helps you with security and also performance improvement in some cases.
	//you can read more about it here.
	// stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4: execution of our sql statement.
	// res, err := stmt.Exec(link.Title, link.Address)
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}
	//#5: retrieving Id of inserted Link.
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

// Query All
func GetAll() []Link {
	// stmt, err := database.Db.Prepare("select id, title, address from Links")
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID") // changed

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		// err := rows.Scan(&link.ID, &link.Title, &link.Address)
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		} // changed
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
