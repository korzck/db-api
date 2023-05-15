package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

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
);`

func StorageDELETE(id int) {
	query := "DELETE FROM storages WHERE id = " + strconv.Itoa(id)
	_, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return
	}
}

type StorageList struct {
	Storages []Storage `json:"storages"`
}

func StorageGET(id int) (Storage, error) {
	query := "SELECT * FROM Storages WHERE id = " + strconv.Itoa(id)
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return Storage{}, err
	}
	if !rows.Next() {
		return Storage{}, errors.New("bad request")
	}
	res := ParseStorageRow(rows)
	if res.Id == 0 {
		return Storage{}, errors.New("bad request")
	}
	return res, nil
} 

func StoragesGET() (StorageList, error) {
	res := new(StorageList)
	query := "SELECT * FROM Storages;"
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		res.Storages = append(res.Storages, ParseStorageRow(rows))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return *res, nil
}

func StoragePOST(c Storage) error {
	sqlLine := "INSERT INTO Storages(name, phone) VALUES "
	str := "( '" + c.Name + "', '" + c.Phone  + "')"
	_, err := DB.Exec(sqlLine + str)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return err
	}
	return nil
}

func ParseStorageRow(rows *sql.Rows) (res Storage) {
	res = Storage{}
	var id int
	var name []byte
	var phone []byte
	if err := rows.Scan(&id, &name, &phone); err != nil {
		log.Fatal("Couldnt scan row ", err)
	}
	nameStr := string(name[:])
	phoneStr := string(phone[:])

	res = Storage{
		Id: id,
		Name: nameStr,
		Phone: phoneStr,
	}
	return
}