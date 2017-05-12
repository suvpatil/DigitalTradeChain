package main

import (
 "github.com/hyperledger/fabric/core/chaincode/shim"
 "errors"
 "encoding/json"
 "math/rand"
 "strconv"
 "time"
)	

func initializeChaincode(stub shim.ChaincodeStubInterface, args []string) error {
	return createDatabase(stub, args)
}

func SaveDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var contractDetails contract
	var err error
	var ok bool
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Need 1 arguments")
	}
	
	contractsId := strconv.Itoa(rand.Int())
	json.Unmarshal([]byte(args[0]), &contractDetails)

	current_time := time.Now().Local()
	contractDetails.contractCreateDate = current_time.Format("2006-01-02")
	contractDetails.isLCAttached = false
	contractDetails.isPOAttached = true
	contractDetails.isBillOfLedingAttached = false
	contractDetails.isInvoiceListAttached = false
	contractDetails.actionPendingOn = "Buyer"
	contractDetails.contractStatus = "Contract Created"
	

	ok, err = insertContractDetails(stub, contractsId, contractDetails)
	if !ok && err == nil {
		return nil, errors.New("Error in adding OrderDetails record")
	}

	return nil, err
}

func getContractDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	contractId := args[0]
	contractList, _ := getContractSpecificList(stub, contractId)

	jsonAsBytes, _ := json.Marshal(contractList)
	return jsonAsBytes, nil

}

/*func GetContract(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

}

func UpdateContractStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

}*/
