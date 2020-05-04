package orders

import (
	"encoding/json"
	"fmt"
	"github.com/m3tam3re/errors"
	"github.com/m3tam3re/real/supplierapi"
	"io/ioutil"
	"strconv"
	"time"
)

const path errors.Path = "github.com/m3tam3re/real/supplierapi/orders"

func GetOpen() ([]*ROrder, error) {
	const op errors.Op = "orders.go|func GetOpen()"

	endpoint := "orders?open=true"
	//endpoint := "orders?limit=100&page=1"
	resp, err := supplierapi.StartRequest("GET", endpoint, nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, errors.E(errors.Internal, path, op, err, "error executing http request")
	}
	if resp.StatusCode != 200 {
		return nil, errors.E(errors.Internal, path, op, fmt.Sprintf("statuscode should be 200, got %v", resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	var orders []*ROrder
	err = json.Unmarshal(body, &orders)
	return orders, nil
}

func GetOrder(id string) (*ROrder, error) {
	const op errors.Op = "orders.go|func GetOrder()"

	endpoint := "orders/" + id
	resp, err := supplierapi.StartRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.E(errors.Internal, path, op, err, "error executing http request")
	}
	if resp.StatusCode != 200 {
		return nil, errors.E(errors.Internal, path, op, fmt.Sprintf("statuscode should be 200, got %v", resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	var order ROrder
	err = json.Unmarshal(body, &order)
	return &order, nil
}

// Send() will confirm an order to the REAL api.
func (o *ROrder) Confirm() error {
	const op errors.Op = "orders.go|method Confirm()"

	endpoint := "orders/" + strconv.Itoa(int(o.FulfilmentOrderId)) + "/confirm"
	resp, err := supplierapi.StartRequest("POST", endpoint, nil)
	defer resp.Body.Close()
	if err != nil {
		return errors.E(errors.Internal, path, op, err, "error executing http reuquest")
	}
	if resp.StatusCode != 204 {
		return errors.E(errors.Internal, path, op, fmt.Sprintf("statuscode should be 204, got %v", resp.StatusCode))
	}
	return nil
}

// Send() will post all shipment data for the units of an order to the REAL api. Please note that the Shipmentdata must
// contain valid data.
func (o *ROrder) Send() error {
	const op errors.Op = "orders.go|method Send()"
	for _, u := range o.Units {
		endpoint := "/order-units/" + strconv.Itoa(int(u.IdOrderUnit)) + "/send"
		body, _ := json.Marshal(u.ShipmentData)
		resp, err := supplierapi.StartRequest("POST", endpoint, body)
		time.Sleep(time.Second * 1)
		if err != nil {
			return errors.E(errors.Internal, path, op, err, "error executing request")
		}
		if resp.StatusCode != 204 {
			return errors.E(errors.Internal, path, op, fmt.Sprintf("statuscode should be 204, got %v", resp.StatusCode))
		}
	}
	return nil
}
