package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Ilaion/lostor/internal/models"
)

type App struct {
	series   models.SeriesModel
	torrents models.TorrentModel
}

func main() {
	db, err := sql.Open("sqlite3", "lostor.db")
	if err != nil {
		log.Fatalln(err)
	}

	app := &App{
		series:   models.SeriesModel{DB: db},
		torrents: models.TorrentModel{DB: db},
	}

	series, err := app.series.All()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(series)
}
