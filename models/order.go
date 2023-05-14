package models

type Order struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Weight float32 `json:"weight"`
	Storage int `json:"storage"`
}

var OrderSQL = `
CREATE TABLE orders (
	id SERIAL,
	date TEXT,
	courier INT,
	client INT,
	address TEXT,
	PRIMARY KEY (id),
	CONSTRAINT fk_courier FOREIGN KEY(courier) REFERENCES couriers(id),
	CONSTRAINT fk_client FOREIGN KEY(client) REFERENCES clients(id)
);
`