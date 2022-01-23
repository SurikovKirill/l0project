package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`          //
	Entry             string    `json:"entry"`              //
	InternalSignature string    `json:"internal_signature"` //
	Payment           Payment   `json:"payment"`            //
	Items             []Item    `json:"items"`              //
	Locale            string    `json:"locale"`             //
	CustomerID        string    `json:"customer_id"`        //
	Delivery          Delivery  `json:"delivery"`           //
	DeliveryService   string    `json:"delivery_service"`   //
	ShardKey          string    `json:"shardkey"`           //
	SmID              int       `json:"sm_id"`              //
	DateCreated       time.Time `json:"date_created"`       //
	OofShard          string    `json:"oof_shard"`          //
}

func (o *Order) Validate() error {
	return validation.ValidateStruct(o)
}
