package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`          //
	Entry             string    `json:"entry"`              //
	InternalSignature string    `json:"internal_signature"` //
	Payment           Payment   `json:"payment"`            //
	Items             []Items   `json:"items"`              //
	Locale            string    `json:"locale"`             //
	CustomerID        string    `json:"customer_id"`        //
	Delivery          Delivery  `json:"delivery"`           //
	DeliveryService   string    `json:"delivery_service"`   //
	ShardKey          string    `json:"shardkey"`           //
	SmID              int       `json:"sm_id"`              //
	DateCreated       time.Time `json:"date_created"`       //
	OofShard          string    `json:"oof_shard"`          //
}
type Payment struct {
	Transaction  string `json:"transaction"`   //
	Currency     string `json:"currency"`      //
	Provider     string `json:"provider"`      //
	Amount       int    `json:"amount"`        //
	PaymentDt    int    `json:"payment_dt"`    //
	Bank         string `json:"bank"`          //
	DeliveryCost int    `json:"delivery_cost"` //
	GoodsTotal   int    `json:"goods_total"`   //
	RequestId    string `json:"request_id"`    //
	CustomFee    int    `json:"custom_fee"`    //
}
type Items struct {
	ChrtID      int    `json:"chrt_id"`      //
	Price       int    `json:"price"`        //
	Rid         string `json:"rid"`          //
	Name        string `json:"name"`         //
	Sale        int    `json:"sale"`         //
	Size        string `json:"size"`         //
	TotalPrice  int    `json:"total_price"`  //
	NmID        int    `json:"nm_id"`        //
	Brand       string `json:"brand"`        //
	Status      int    `json:"status"`       //
	TrackNumber string `json:"track_number"` //
}

type Delivery struct {
	Name    string `json:"name"`    //
	Phone   string `json:"phone"`   //
	Zip     string `json:"zip"`     //
	City    string `json:"city"`    //
	Address string `json:"address"` //
	Region  string `json:"region"`  //
	Email   string `json:"email"`   //
}

const (
	clusterID = "test-cluster"
	clientID  = "test-pub"
)

func main() {
	fmt.Println("Connecting")
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4223"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Publishing")
	// Simple Synchronous Publisher
	d := &Delivery{
		Address: "Ploshad Mira 15",
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}
	i := &Items{
		ChrtID:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmID:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}
	var it []Items
	it = append(it, *i)
	p := &Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestId:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}

	o := &Order{
		OrderUID:          "b563feb7b2b84b6test",
		Entry:             "WBIL",
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "",
		Delivery:          *d,
		Payment:           *p,
		Items:             it,
	}

	out, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
	sc.Publish("test", out)
	sc.Close()
}
