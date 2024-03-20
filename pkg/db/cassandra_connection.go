package db

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
	session *gocql.Session
)

type Person struct {
	ID      gocql.UUID `json: ID`
	Name    string     `json: name`
	Surname string     `json: surname`
}

func init() {
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "my_keyspace"
	cluster.Consistency = gocql.Quorum

	// Create conn
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	// Create table
	if err := session.Query(`CREATE TABLE IF NOT EXISTS person (id UUID PRIMARY KEY, name TEXT, surname TEXT)`).Exec(); err != nil {
		panic(err)
	}
}

// Create
func Create(name string, age int) error {
	id := gocql.TimeUUID()
	return session.Query(`INSERT INTO person (id, name, age) VALUES (?, ?, ?)`, id, name, age).Exec()
}

// Read
func Get(id gocql.UUID) (*Person, error) {
	var person Person
	if err := session.Query(`SELECT id, name, age FROM person WHERE id = ?`, id).Scan(&person.ID, &person.Name, &person.Surname); err != nil {
		return nil, err
	}
	return &person, nil
}

// Update
func Update(id gocql.UUID, name, surname string) error {
	return session.Query(`UPDATE person SET name = ?, surname = ? WHERE id = ?`, name, surname, id).Exec()
}

// Delete
func Delete(id gocql.UUID) error {
	return session.Query(`DELETE FROM person WHERE id = ?`, id).Exec()
}
