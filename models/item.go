package models

type Item struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Weight float32 `json:"weight"`
	Storage int `json:"storage"`
}

var ItemSQL = `
CREATE TABLE items (
	id SERIAL,
	name TEXT,
	description TEXT,
	weight REAL,
	storage INT,
	PRIMARY KEY (id),
	CONSTRAINT fk_storage FOREIGN KEY(storage) REFERENCES storages(id)
);
`