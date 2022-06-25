CREATE TABLE series (
  id INTEGER,
  name VARCHAR(64) NOT NULL UNIQUE,
  quality VARCHAR(10) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE torrents (
  id INTEGER,
  season INTEGER NOT NULL,
  episode INTEGER NOT NULL,
  created_at DATETIME NOT NULL,
  series_id INTEGER,
  PRIMARY KEY (id),
  FOREIGN KEY (series_id) REFERENCES series(id)
);
