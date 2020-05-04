package orders

import (
	"testing"
)

func TestGetOpen(t *testing.T) {
	orders, err := GetOpen()
	if err != nil {
		t.Error(err)
	}
	t.Log(orders)
	//TODO add tests
}

func TestGetOrder(t *testing.T) {
	order, err := GetOrder("6025488")
	if err != nil {
		t.Error(err)
	}
	order.Send()
	t.Log(order)
}
