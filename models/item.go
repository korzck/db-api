package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

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
);`

func ItemDELETE(id int) {
	query := "DELETE FROM items WHERE id = " + strconv.Itoa(id)
	_, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return
	}
}

type ItemList struct {
	Items []Item `json:"items"`
}

func ItemGET(id int) (Item, error) {
	query := "SELECT * FROM items WHERE id = " + strconv.Itoa(id)
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return Item{}, err
	}
	if !rows.Next() {
		return Item{}, errors.New("bad request")
	}
	res := ParseItemRow(rows)
	if res.Id == 0 {
		return Item{}, errors.New("bad request")
	}
	return res, nil
} 

func ItemsGET() (ItemList, error) {
	res := new(ItemList)
	query := "SELECT * FROM items;"
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		res.Items = append(res.Items, ParseItemRow(rows))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return *res, nil
}

func ItemPOST(c Item) error {
	sqlLine := "INSERT INTO Items(name, description, weight, storage) VALUES "
	str := "( '" + c.Name + "', '" + c.Description + "'," + fmt.Sprintf("%f", c.Weight) + ", " + strconv.Itoa(c.Storage) + ")"
	_, err := DB.Exec(sqlLine + str)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return err
	}
	return nil
}

func ParseItemRow(rows *sql.Rows) (res Item) {
	res = Item{}
	var id int
	var name []byte
	var description []byte
	var weight float32
	var storage int
	if err := rows.Scan(&id, &name, &description, &weight, &storage); err != nil {
		log.Fatal("Couldnt scan row ", err)
	}
	nameStr := string(name[:])
	descriptionStr := string(description[:])
	res = Item{
		Id: id,
		Name: nameStr,
		Description: descriptionStr,
		Weight: weight,
		Storage: storage,
	}
	return
}