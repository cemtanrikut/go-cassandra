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
