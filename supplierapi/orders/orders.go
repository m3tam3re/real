package orders

import (
	"encoding/json"
	"fmt"
	"github.com/m3tam3re/errors"
	"github.com/m3tam3re/real/supplierapi"
	"io/ioutil"
)

const path errors.Path = "github.com/m3tam3re/real/supplierapi/orders"

func GetOpen() ([]Order, error) {
	const op errors.Op = "orders.go|func GetOpen()"

	//endpoint := "orders?open=true"
	endpoint := "orders?limit=100&page=1"
	resp, err := supplierapi.StartRequest("GET", endpoint, nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, errors.E(errors.Internal, path, op, err, "error executing http request")
	}
	if resp.StatusCode != 200 {
		return nil, errors.E(errors.Internal, path, op, err, fmt.Sprintf("statuscode should be 200, got %v", resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	var orders []Order
	err = json.Unmarshal(body, &orders)
	return orders, nil
}

func (o *Order) Confirm() error {
	const op errors.Op = "orders.go|method Confirm()"

	endpoint := "/orders/{id}/confirm" + o.FulfilmentOrderId + "/confirm"
	resp, err := supplierapi.StartRequest("POST", endpoint, nil)
	defer resp.Body.Close()
	if err != nil {
		return errors.E(errors.Internal, path, op, err, "error executing http reuquest")
	}
	if resp.StatusCode != 204 {
		return errors.E(errors.Internal, path, op, err, fmt.Sprintf("statuscode should be 204, got %v", resp.StatusCode))
	}
	return nil
}
