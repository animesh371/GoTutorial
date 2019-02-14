package ecomplatform

import (
	"fmt"
	"testing"
)

func TestOrderManager(t *testing.T) {
	orderManager := &OrderManager{orders: make(map[string]Order)}
	items := []item{{10, "bread"}, {20, "milk"}, {10, "biscuit"}}
	requests := []request{{1, items[0]}, {2, items[2]}, {5, items[1]}, {6, items[0]}}
	orderIds := []string{}
	requestID := orderManager.createNewOrder(requests[0])

	orderIds = append(orderIds, requestID)
	pendingOrders := orderManager.getPendingOrders()

	fmt.Println(len(pendingOrders))
}
