package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Order struct {
	Id int `json:"id"`
	Date string `json:"date"`
	Courier int `json:"courier"`
	Client int `json:"client"`
	Address string `json:"address"`
}

var OrderSQL = `
CREATE TABLE orders (
	id SERIAL,
	date DATE,
	courier INT,
	client INT,
	address TEXT,
	PRIMARY KEY (id),
	CONSTRAINT fk_courier FOREIGN KEY(courier) REFERENCES couriers(id),
	CONSTRAINT fk_Order FOREIGN KEY(Order) REFERENCES Orders(id)
);
`

type OrderList struct {
	Orders []Order `json:"orders"`
}

func OrderDELETE(id int) {
	query := "DELETE FROM orders WHERE id = " + strconv.Itoa(id)
	_, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return
	}
}

func OrderGET(id int) (Order, error) {
	query := "SELECT * FROM orders WHERE id = " + strconv.Itoa(id)
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal("Couldnt query", err)
		return Order{}, err
	}
	if !rows.Next() {
		return Order{}, errors.New("bad request")
	}
	res := ParseOrderRow(rows)
	if res.Id == 0 {
		return Order{}, errors.New("bad request")
	}
	return res, nil
} 

func OrdersGET() (OrderList, error) {
	res := new(OrderList)
	query := "SELECT * FROM Orders;"
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		res.Orders = append(res.Orders, ParseOrderRow(rows))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return *res, nil
}

func OrderPOST(c Order) error {
	sqlLine := "INSERT INTO orders(date, courier, client, address) VALUES "
	date := strings.Split(c.Date, "/")
	dateStr := fmt.Sprintf("TO_DATE('%v/%v/%v', 'DD/MM/YYYY')", date[0], date[1], date[2])
	str := "(" + dateStr + ", '" + strconv.Itoa(c.Courier) + ", " + strconv.Itoa(c.Client) + ", '" +c.Address + "')"
	_, err := DB.Exec(sqlLine + str)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return err
	}
	return nil
}

func ParseOrderRow(rows *sql.Rows) (res Order) {
	res = Order{}
	var id int
	var name []byte
	var phone []byte
	if err := rows.Scan(&id, &date, &courier); err != nil {
		log.Fatal("Couldnt scan row ", err)
	}
	nameStr := string(name[:])
	phoneStr := string(phone[:])

	res = Order
	return
}