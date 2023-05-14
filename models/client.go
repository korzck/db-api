package models

type Client struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Weight float32 `json:"weight"`
	Storage int `json:"storage"`
}

var ClientSQL = `
CREATE TABLE clients (
	id SERIAL,
	name TEXT,
	phone TEXT,
	PRIMARY KEY (id)
);
`