package store

import (
	"encoding/json"
	"l0project/internal/model"
	"log"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(o *model.Order) (*model.Order, error) {
	or, err := json.Marshal(o)
	if err != nil {
		log.Println(err)
	}
	log.Println("Marshalled object")
	log.Println(string(or))
	if err := r.store.db.QueryRow("INSERT INTO orders (order_uid, info) values ($1, $2) returning order_uid", o.OrderUID, or).Scan(&o.OrderUID); err != nil {
		log.Println(err)
		return nil, err
	}

	return o, nil
}

func (r *OrderRepository) AllOrders() ([]model.Order, error) {
	log.Println("Select all orders from db for cache")
	tmp, err := r.store.db.Query("SELECT * from orders")
	if err != nil {
		return nil, err
	}
	o := model.Order{}
	var result []model.Order
	var info []byte
	var p string
	log.Println("Prepare data from db for cache")
	for tmp.Next() {
		if err := tmp.Scan(&p, &info); err != nil {
			log.Println(err)
		}
		err := json.Unmarshal(info, &o)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Row from db")
		log.Println(o)
		result = append(result, o)
	}
	log.Println("Result for cache")
	log.Println(result)
	return result, nil
}
