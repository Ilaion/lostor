package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	DB *sql.DB
}

type Series struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Quality string `json:"quality"`
}

type Torrent struct {
	ID        int `json:"id"`
	Season    int `json:"season"`
	Episode   int `json:"episode"`
	CreatedAt int `json:"created_at"`
	SeriesID  int `json:"series_id"`
}

func main() {
	db, err := sql.Open("sqlite3", "lostor.db")
	if err != nil {
		log.Fatalln(err)
	}

	app := &Application{
		DB: db,
	}

	// id, err := app.addSeries(Series{
	// 	Name:    "Peacky Blinders",
	// 	Quality: "1080p",
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(id)

	// err = app.deleteSeries(1)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// series, err := app.getSeries()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(series)

	// series, err := app.getSeriesByID(1)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(series)

	// TORRENT
	torrents, err := app.getTorrents()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(torrents)
}

func (app *Application) getSeries() ([]Series, error) {
	rows, err := app.DB.Query("SELECT * FROM series")
	if err != nil {
		return []Series{}, err
	}

	series := []Series{}
	for rows.Next() {
		var s Series

		err = rows.Scan(&s.ID, &s.Name, &s.Quality)
		if err != nil {
			return []Series{}, err
		}

		series = append(series, s)
	}

	return series, nil
}

func (app *Application) getSeriesByID(idx int) (Series, error) {
	row := app.DB.QueryRow("SELECT * FROM series WHERE id=?", idx)

	var series Series

	err := row.Scan(&series.ID, &series.Name, &series.Quality)
	if err != nil {
		return Series{}, err
	}

	return series, nil
}

func (app *Application) addSeries(series Series) (int, error) {
	stmt := "INSERT INTO series VALUES(NULL,?,?)"

	res, err := app.DB.Exec(stmt, series.Name, series.Quality)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (app *Application) deleteSeries(idx int) error {
	stmt := "DELETE FROM series WHERE id=?"

	_, err := app.DB.Exec(stmt, idx)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) getTorrents() ([]Torrent, error) {
	rows, err := app.DB.Query("SELECT * FROM torrents")
	if err != nil {
		return []Torrent{}, err
	}

	torrents := []Torrent{}
	for rows.Next() {
		var t Torrent

		err = rows.Scan(&t.ID, &t.Season, &t.Episode, &t.CreatedAt, &t.SeriesID)
		if err != nil {
			return []Torrent{}, err
		}

		torrents = append(torrents, t)
	}

	return torrents, nil
}

func (app *Application) addTorrent(t Torrent) (int, error) {
	stmt := "INSERT INTO torrents VALUES(NULL,?,?,?,?)"

	res, err := app.DB.Exec(stmt, t.Season, t.Episode, t.CreatedAt, t.SeriesID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (app *Application) deleteTorrent(idx int) error {
	stmt := "DELETE FROM torrents WHERE id=?"

	_, err := app.DB.Exec(stmt, idx)
	if err != nil {
		return err
	}

	return nil
}
