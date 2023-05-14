package models

type Courier struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Weight float32 `json:"weight"`
	Storage int `json:"storage"`
}

var CourierSQL = `
CREATE TABLE couriers (
	id SERIAL,
	name TEXT,
	phone TEXT,
	car TEXT,
	PRIMARY KEY (id)
);
`