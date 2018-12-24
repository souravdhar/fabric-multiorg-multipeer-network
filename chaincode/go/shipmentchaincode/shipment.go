package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define simple contract structure
type shipment struct {
}

// Shipment ...
type Shipment struct {
	OrderID                string `json:"orderid"`
	ShipmentRequestDate    string `json:"shipmentrequestdate"`
	ShipmentFrom           string `json:"shipmentfrom"`
	ShipmentTo             string `json:"shipmentto"`
	ShipmentCompletionDate string `json:"shipmentcompletiondate"`
	ShipmentQuantity       int    `json:"shipmentquantity"`
}

func (s *shipment) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *shipment) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Simple Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	//Route to appropriate handler function
	if function == "queryShipment" {
		return s.queryShipment(APIstub, args)
	} else if function == "createShipment" {
		return s.createShipment(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	}

	return shim.Error("Invalid smart contract function name.")
}

func (s *shipment) queryShipment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		shim.Error("Incorrect number of arguments. Expecting 1")
	}
	shipmentDetailsAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(shipmentDetailsAsBytes)
}

func (s *shipment) createShipment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	quantity, _ := strconv.Atoi(args[6])
	shipmentDetails := Shipment{OrderID: args[1], ShipmentRequestDate: args[2], ShipmentFrom: args[3], ShipmentTo: args[4], ShipmentCompletionDate: args[5], ShipmentQuantity: quantity}

	shipmentDetailsAsBytes, _ := json.Marshal(shipmentDetails)
	APIstub.PutState(args[0], shipmentDetailsAsBytes)

	return shim.Success(nil)
}

func (s *shipment) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	order := []Shipment{
		Shipment{OrderID: "O1", ShipmentRequestDate: "2018-12-24", ShipmentFrom: "S1", ShipmentTo: "B1", ShipmentCompletionDate: "", ShipmentQuantity: 10000},
		Shipment{OrderID: "O2", ShipmentRequestDate: "2018-12-24", ShipmentFrom: "S1", ShipmentTo: "B2", ShipmentCompletionDate: "", ShipmentQuantity: 12000},
	}

	i := 1
	for i < len(order) {
		fmt.Println("i is ", i)
		orderAsBytes, _ := json.Marshal(order[i])
		APIstub.PutState("O"+strconv.Itoa(i), orderAsBytes)
		fmt.Println("Added", order[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(shipment))
	if err != nil {
		fmt.Printf("Error creating new smart contract: %s", err)
	}
}
