package models

type Date string

type OrderStatus string

const (
	PlacedOrderStatus    OrderStatus = "placed"
	ApprovedOrderStatus  OrderStatus = "approved"
	DeliveredOrderStatus OrderStatus = "delivered"
)

type Order struct {
	Id       int64       `json:"id"`
	PetId    int64       `json:"petId"`
	Quantity int32       `json:"quantity"`
	ShipDate Date        `json:"shipDate"`
	Status   OrderStatus `json:"status"`
	Complete bool        `json:"complete"`
}
