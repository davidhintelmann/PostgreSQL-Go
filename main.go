package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// user, password, and databasename for postgresql instance
const user, password, dbname = "david", "[Germany92", "AdventureWorks2014"

// use background context globally to pass between functions
var ctx = context.Background()

type People struct {
	Persons []Person `json:"people"`
}

// type Person struct {
// 	ID         int            `json:"id"`
// 	Title      sql.NullString `json:"title"`
// 	FirstName  string         `json:"firstname"`
// 	MiddleName sql.NullString `json:"middlename"`
// 	LastName   string         `json:"lastname"`
// 	Suffix     sql.NullString `json:"suffix"`
// 	Scode      string         `json:"scode"`
// 	Ccode      string         `json:"ccode"`
// 	State      string         `json:"state"`
// 	Country    string         `json:"country"`
// }

type Person struct {
	BID          int            `json:"businessid"`
	PersonType   string         `json:"persontype"`
	NameStyle    bool           `json:"namestyle"`
	Title        sql.NullString `json:"title"`
	FirstName    string         `json:"firstname"`
	MiddleName   sql.NullString `json:"middlename"`
	LastName     string         `json:"lastname"`
	Suffix       sql.NullString `json:"suffix"`
	EmailPromo   int            `json:"emailpromo"`
	ContactInfo  sql.NullString `json:"contactinfo"`
	Demographics string         `json:"demographics"`
	UUID         string         `json:"uuid"`
	Timestamp    time.Time      `json:"timestamp"`
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	fmt.Print("Connecting to postgresql...\n")
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	dbpool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatalf("error during intial connection: %v\n", err)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("error during QueryRow: %v\n", err)
	}
	fmt.Println(greeting)

	people := People{}
	query := fmt.Sprintf(`SELECT * FROM person.person`)

	rows, err := dbpool.Query(ctx, query)
	if err != nil {
		log.Fatalf("error during Query: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var person Person
		err = rows.Scan(
			&person.BID,
			&person.PersonType,
			&person.NameStyle,
			&person.Title,
			&person.FirstName,
			&person.MiddleName,
			&person.LastName,
			&person.Suffix,
			&person.EmailPromo,
			&person.ContactInfo,
			&person.Demographics,
			&person.UUID,
			&person.Timestamp,
		)

		if err != nil {
			log.Fatalf("error during rows.Scan(): %v\n", err)
		}
		people.AddPerson(person)
	}
	// fmt.Println(people)
	fmt.Println(people.getCount())
}

func (peeps *People) AddPerson(per Person) {
	peeps.Persons = append(peeps.Persons, per)
}

func (peeps *People) getCount() int {
	return len(peeps.Persons)
}
