package ecomplatform

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type item struct {
	price       int
	description string
}

type request struct {
	quantity    int
	orderedItem item
}
type amend struct {
	quantity int
	orderID  string
}

type Order struct {
	quantity    int
	orderedItem item
	status      string
}
type orderHandling interface {
	createNewOrder(req request) string
	getOrderStatus(id string) string
	modifyOrder(replace amend) string
	updateStatus(id string, status string) string
	getPendingOrders() []Order
}

type OrderManager struct {
	// Order Manager
	orders map[string]Order
}

func (orderManager *OrderManager) createNewOrder(req request) string {
	orderID := randSeq(10)
	order := Order{req.quantity, req.orderedItem, "Pending"}
	orderManager.orders[orderID] = order
	return orderID
}

func (orderManager *OrderManager) getOrderStatus(id string) string {
	elem, ok := orderManager.orders[id]
	if ok {
		return elem.status
	}
	return "INVALID"
}
func (orderManager *OrderManager) modifyOrder(replace amend) string {

	if elem, ok := orderManager.orders[replace.orderID]; ok {
		orderManager.orders[replace.orderID] = Order{replace.quantity, elem.orderedItem, elem.status}
		return replace.orderID
	}
	return "INVALID"
}

func (orderManager *OrderManager) updateStatus(id string, status string) string {
	if elem, ok := orderManager.orders[id]; ok {
		orderManager.orders[id] = Order{elem.quantity, elem.orderedItem, status}
		return id
	}
	return "INVALID"
}

func (orderManager *OrderManager) getPendingOrders() []Order {
	var pendingOrders []Order
	pendingStatus := "PENDING"
	for orderID, order := range orderManager.orders {
		_ = orderID
		if strings.Compare(order.status, pendingStatus) == 0 {
			pendingOrders = append(pendingOrders, order)
		}
	}
	return pendingOrders
}
