package orders

import (
	"encoding/json"
	"fmt"
	"github.com/m3tam3re/errors"
	"github.com/m3tam3re/real/supplierapi"
	"io/ioutil"
	"strconv"
)

const path errors.Path = "github.com/m3tam3re/real/supplierapi/orders"

func GetOpen() ([]ROrder, error) {
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
	var orders []ROrder
	err = json.Unmarshal(body, &orders)
	return orders, nil
}

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

func (o *ROrder) Send() error {
	/*const op errors.Op = "orders.go|method Send()"

	var body [][]byte
	for _, u := range o.Units {
		sd, _ := json.Marshal(u.ShipmentData)
		body = append(body, sd)
	}

	/*endpoint := "/order-units/" + strconv.Itoa(int(o.FulfilmentOrderId)) + "/send"
	resp, err := supplierapi.StartRequest("POST", endpoint, body)
	defer resp.Body.Close()
	if err != nil {
		return errors.E(errors.Internal, path, op, err, "error executing request")
	}*/
	return nil
}
