package models

import "database/sql"

type Torrent struct {
	ID        int `json:"id"`
	Season    int `json:"season"`
	Episode   int `json:"episode"`
	CreatedAt int `json:"created_at"`
	SeriesID  int `json:"series_id"`
}

type TorrentModel struct {
	DB *sql.DB
}

func (tm *TorrentModel) All() ([]Torrent, error) {
	rows, err := tm.DB.Query("SELECT * FROM torrents")
	if err != nil {
		return []Torrent{}, err
	}
	defer rows.Close()

	var ts []Torrent

	for rows.Next() {
		var t Torrent

		err = rows.Scan(&t.ID, &t.Season, &t.Episode, &t.CreatedAt, &t.SeriesID)
		if err != nil {
			return nil, err
		}

		ts = append(ts, t)
	}

	return ts, nil
}

func (tm *TorrentModel) Add(t Torrent) (int, error) {
	res, err := tm.DB.Exec("INSERT INTO torrents VALUES(NULL,?,?,?,?)", t.Season, t.Episode, t.CreatedAt, t.SeriesID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tm *TorrentModel) Delete(id int) error {
	_, err := tm.DB.Exec("DELETE FROM torrents WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}
