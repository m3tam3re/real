package orders

import (
	"testing"
)

func TestGetOpen(t *testing.T) {
	orders, err := GetOpen()
	if err != nil {
		t.Error(err)
	}
	for _, order := range orders {
		t.Log(order)
	}
}
