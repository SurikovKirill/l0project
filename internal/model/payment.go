package model

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
