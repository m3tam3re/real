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
