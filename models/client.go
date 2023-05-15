package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Client struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

var ClientSQL = `
CREATE TABLE clients (
	id SERIAL,
	name TEXT,
	phone TEXT,
	PRIMARY KEY (id)
);
`

type ClientList struct {
	Clients []Client `json:"clients"`
}

func ClientDELETE(id int) {
	query := "DELETE FROM clients WHERE id = " + strconv.Itoa(id)
	_, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return
	}
}

func ClientGET(id int) (Client, error) {
	query := "SELECT * FROM Clients WHERE id = " + strconv.Itoa(id)
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return Client{}, err
	}
	if !rows.Next() {
		return Client{}, errors.New("bad request")
	}
	res := ParseClientRow(rows)
	if res.Id == 0 {
		return Client{}, errors.New("bad request")
	}
	return res, nil
} 

func ClientsGET() (ClientList, error) {
	res := new(ClientList)
	query := "SELECT * FROM Clients;"
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		res.Clients = append(res.Clients, ParseClientRow(rows))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return *res, nil
}

func ClientPOST(c Client) error {
	sqlLine := "INSERT INTO Clients(name, phone) VALUES "
	str := "( '" + c.Name + "', '" + c.Phone  + "')"
	_, err := DB.Exec(sqlLine + str)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return err
	}
	return nil
}

func ParseClientRow(rows *sql.Rows) (res Client) {
	res = Client{}
	var id int
	var name []byte
	var phone []byte
	if err := rows.Scan(&id, &name, &phone); err != nil {
		log.Fatal("Couldnt scan row ", err)
	}
	nameStr := string(name[:])
	phoneStr := string(phone[:])

	res = Client{
		Id: id,
		Name: nameStr,
		Phone: phoneStr,
	}
	return
}