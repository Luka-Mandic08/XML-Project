package create_order

type Color struct {
	Code string
}

type Product struct {
	Id    string
	Color Color
}

type OrderItem struct {
	Product  Product
	Quantity uint16
}

type OrderDetails struct {
	Id      string
	Items   []OrderItem
	Address string
}

type CreateOrderCommandType int8

const (
	UpdateInventory CreateOrderCommandType = iota
	RollbackInventory
	ApproveOrder
	CancelOrder
	ShipOrder
	UnknownCommand
)

type CreateOrderCommand struct {
	Order OrderDetails
	Type  CreateOrderCommandType
}

type CreateOrderReplyType int8

const (
	InventoryUpdated CreateOrderReplyType = iota
	InventoryNotUpdated
	InventoryRolledBack
	OrderShippingScheduled
	OrderShippingNotScheduled
	OrderApproved
	OrderCancelled
	UnknownReply
)

type CreateOrderReply struct {
	Order OrderDetails
	Type  CreateOrderReplyType
}
