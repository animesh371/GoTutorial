package ecomplatform

import (
	"strings"
	"testing"
)

func TestOrderManager(t *testing.T) {
	orderManager := &OrderManager{orders: make(map[string]Order)}
	items := []item{{10, "bread"}, {20, "milk"}, {10, "biscuit"}}
	requests := []request{{1, items[0]}, {2, items[2]}, {5, items[1]}, {6, items[0]}}
	requestIds := []string{}
	orderIds := []string{}
	for index, request := range requests {
		_ = index
		requestID := orderManager.createNewOrder(request)
		requestIds = append(requestIds, requestID)
		orderIds = append(orderIds, requestID)
	}

	pendingOrders := orderManager.getPendingOrders()

	if len(pendingOrders) != 4 {
		t.Errorf("Expected 4 pending orders, actual pending orders %d", len(pendingOrders))
	}

	amendOrder := amend{4, requestIds[1]}
	orderManager.modifyOrder(amendOrder)

	orderManager.updateStatus(requestIds[1], "Dispatched")
	orderManager.updateStatus(requestIds[2], "Cancelled")
	pendingOrders = orderManager.getPendingOrders()

	if len(pendingOrders) != 2 {
		t.Errorf("Expected 4 pending orders, actual pending orders %d", len(pendingOrders))
	}

	if strings.Compare(orderManager.getOrderStatus(requestIds[1]), "Dispatched") != 0 {
		t.Error("Order status is", orderManager.getOrderStatus(requestIds[1]))
	}

	if strings.Compare(orderManager.getOrderStatus(requestIds[2]), "Cancelled") != 0 {
		t.Error("Order status is", orderManager.getOrderStatus(requestIds[1]))
	}
}
