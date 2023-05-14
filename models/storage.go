package models

type Storage struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

var StorageSQL = `
CREATE TABLE storages (
	id SERIAL,
	name TEXT,
	phone TEXT,
	PRIMARY KEY (id)
);
`