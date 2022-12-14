package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	Id       int    ` bson:"id" json:"id" `
	Category string ` bson:"category" json:"category"`
	Img      string ` bson:"img" json:"img"`
	Title    string ` bson:"title" json:"title"`
	Price    string ` bson:"price" json:"price"`
	Count    int    ` bson:"count" json:"count"`
}

type OrderData struct {
	Adress string `bson:"adress" json:"adress"`
	Mobile string `bson:"mobile" json:"mobile"`
	Time   string `bson:"time" json:"time"`
}

type Order struct {
	OrderId   int `json:"order_id" bson:"order_id"`
	Products  []Product
	Data      OrderData
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type DBResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	OrderId   int                `json:"order_id" bson:"order_id"`
	Products  []Product
	Data      OrderData
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (o *Order) String() string {
	data := fmt.Sprintf("Номер заказа - <b>%v</b>\n%v", o.OrderId, o.Data.String())
	data += "<b>Блюда</b>:\n"
	for _, product := range o.Products {
		data += product.String()
	}
	return data
}
func (p *Product) String() string {
	return fmt.Sprintf("%v - %v\nВсего %v ед. - стоимость %vр. за ед.\n******************\n", p.Category, p.Title, p.Count, p.Price)
}
func (o *OrderData) String() string {
	return fmt.Sprintf("<b>Данные по заказу:</b>\nАдрес - %v\nНомер - %v\nВремя доставки - %v\n", o.Adress, o.Mobile, o.Time)
}
