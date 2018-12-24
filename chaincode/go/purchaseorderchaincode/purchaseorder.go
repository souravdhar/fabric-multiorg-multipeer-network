package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define simple contract structure
type purchaseorder struct {
}

// PurchaseOrder ...
type PurchaseOrder struct {
	OrderID       string `json:"orderid"`
	OrderDate     string `json:"orderdate"`
	BuyerID       string `json:"buyerid"`
	SupplierID    string `json:"supplierid"`
	OrderQuantity int    `json:"orderquantity"`
}

func (s *purchaseorder) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *purchaseorder) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Simple Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	//Route to appropriate handler function
	if function == "queryOrder" {
		return s.queryOrder(APIstub, args)
	} else if function == "createOrder" {
		return s.createOrder(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	}

	return shim.Error("Invalid smart contract function name : " + function)
}

func (s *purchaseorder) queryOrder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		shim.Error("Incorrect number of arguments. Expecting 1")
	}
	orderAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(orderAsBytes)
}

func (s *purchaseorder) createOrder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	quantity, _ := strconv.Atoi(args[5])
	order := PurchaseOrder{OrderID: args[1], OrderDate: args[2], BuyerID: args[3], SupplierID: args[4], OrderQuantity: quantity}

	orderAsBytes, _ := json.Marshal(order)
	APIstub.PutState(args[0], orderAsBytes)

	return shim.Success(nil)
}

func (s *purchaseorder) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	order := []PurchaseOrder{
		PurchaseOrder{OrderID: "O1", OrderDate: "2018-12-24", BuyerID: "B1", SupplierID: "S1", OrderQuantity: 10000},
		PurchaseOrder{OrderID: "O2", OrderDate: "2018-12-24", BuyerID: "B2", SupplierID: "S1", OrderQuantity: 12000},
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
	err := shim.Start(new(purchaseorder))
	if err != nil {
		fmt.Printf("Error creating new smart contract: %s", err)
	}
}
