package models

import "database/sql"

type Series struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Quality string `json:"quality"`
}

type SeriesModel struct {
	DB *sql.DB
}

func (sm *SeriesModel) All() ([]Series, error) {
	rows, err := sm.DB.Query("SELECT * FROM series")
	if err != nil {
		return []Series{}, err
	}
	defer rows.Close()

	var series []Series

	for rows.Next() {
		var s Series

		err = rows.Scan(&s.ID, &s.Name, &s.Quality)
		if err != nil {
			return nil, err
		}

		series = append(series, s)
	}

	return series, nil
}

func (sm *SeriesModel) Get(id int) (Series, error) {
	row := sm.DB.QueryRow("SELECT * FROM series WHERE id=?", id)

	var s Series

	err := row.Scan(&s.ID, &s.Name, &s.Quality)
	if err != nil {
		return Series{}, err
	}

	return s, nil
}

func (sm *SeriesModel) Add(s Series) (int, error) {
	res, err := sm.DB.Exec("INSERT INTO series VALUES(NULL,?,?)", s.Name, s.Quality)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (sm *SeriesModel) Delete(id int) error {
	_, err := sm.DB.Exec("DELETE FROM series WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}
