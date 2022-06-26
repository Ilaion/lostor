package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.GET("/series/", app.getAllSeries)
	router.POST("/series/", app.addSeries)
	router.GET("/series/:id", app.getSeries)
	router.DELETE("/series/:id", app.deleteSeries)
	router.GET("/torrent/", app.getAllTorrents)

	router.Run("localhost:8080")
}
