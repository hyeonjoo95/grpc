package database

import (
	"context"
	"log"
	"main/ent"

	"main/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

var Client *ent.Client

func SetUp() (err error) {
	Client, err = ent.Open("mysql", "testbank:testbank@tcp(127.0.0.1:3306)/user?parseTime=true")
	if err != nil {
		log.Fatalf("failed connecting to MySQL: %v", err)
		return
	}

	err = migration(Client)
	if err != nil {
		log.Println("failed creating schema resources: ", err.Error())
		return
	}

	return
}

func migration(client *ent.Client) error {
	return client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true))
}
