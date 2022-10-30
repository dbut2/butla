package main

import (
	"context"
	"log"

	"github.com/dbut2/shortener-web/pkg/database"
)

func main() {
	s, err := database.New(database.Config{
		Hostname: "localhost:3306",
		Username: "root",
		Database: "shortener",
	})
	if err != nil {
		log.Fatal()
	}

	links, err := s.GetAll(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, link := range links {
		if link.Url == "https://example.com" {
			err = s.Delete(context.Background(), link.Code)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}
