package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Courier struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Car string `json:"car"`
}

type CourierList struct {
	Couriers []Courier `json:"couriers"`
}

var CourierSQL = `
CREATE TABLE couriers (
	id SERIAL,
	name TEXT,
	phone TEXT,
	car TEXT,
	PRIMARY KEY (id)
);`

func CourierDELETE(id int) {
	query := "DELETE FROM couriers WHERE id = " + strconv.Itoa(id)
	_, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return
	}
}

func CourierGET(id int) (Courier, error) {
	query := "SELECT * FROM couriers WHERE id = " + strconv.Itoa(id)
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return Courier{}, err
	}
	if !rows.Next() {
		return Courier{}, errors.New("bad request")
	}
	res := ParseCourierRow(rows)
	if res.Id == 0 {
		return Courier{}, errors.New("bad request")
	}
	return res, nil
} 

func CouriersGET() (CourierList, error) {
	res := new(CourierList)
	query := "SELECT * FROM couriers;"
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		res.Couriers = append(res.Couriers, ParseCourierRow(rows))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return *res, nil
}

func CourierPOST(c Courier) error {
	sqlLine := "INSERT INTO couriers(name, phone, car) VALUES "
	str := "( '" + c.Name + "', '" + c.Phone + "', '" + c.Car + "')"
	_, err := DB.Exec(sqlLine + str)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return err
	}
	return nil
}

func ParseCourierRow(rows *sql.Rows) (res Courier) {
	res = Courier{}
	var id int
	var name []byte
	var phone []byte
	var car []byte
	if err := rows.Scan(&id, &name, &phone, &car); err != nil {
		log.Fatal("Couldnt scan row ", err)
	}
	nameStr := string(name[:])
	phoneStr := string(phone[:])
	carStr := string(car[:])

	res = Courier{
		Id: id,
		Name: nameStr,
		Phone: phoneStr,
		Car: carStr,
	}
	return
}
