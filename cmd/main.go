package main

import (
	"fmt"

	cassandra "github.com/cemtanrikut/go-cassandra/pkg/db"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func main() {
	defer session.Close()

	// CRUD
	// Create
	personID := gocql.TimeUUID()
	read(personID)

	// Read
	p, err := cassandra.Get(personID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Person:", p)
	}

	// Update
	if err := cassandra.Update(personID, "Cem2", "Tanrikut2"); err != nil {
		fmt.Println("Hata:", err)
	}

	//Delete
	read(personID)
	if err := cassandra.Delete(personID); err != nil {
		fmt.Println("Hata:", err)
	}

	// Read again (double check to delete person)
	read(personID)
}

func read(personID gocql.UUID) {

	if err := cassandra.Create(personID, "Cem", "Tanrikut"); err != nil {
		fmt.Println("Error:", err)
	}
}
