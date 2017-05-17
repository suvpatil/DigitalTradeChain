package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Region Chaincode implementation
type DTC_Chaincode struct {
}

func (t *DTC_Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	//Create database on blockchain
	err = initializeChaincode(stub, args)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *DTC_Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "saveContract" {
		// Insert Contract data in blockchain
		return saveContractDetails(stub, args)
	} else if function == "SaveAttachment" {
		// inserting attachment data in blockchain
		return saveAttachmentDetails(stub, args)
	}

	return nil, nil
}

func (t *DTC_Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getContractDetailsByContractId" {
		// Read contract details from blockchain
		return getContractDetailsByContractId(stub, args)
	} else if function == "getAttachment" {
		// get attachment details from blockchain
		return getAttachment(stub, args)
	}else if function == "getContractDetailsByUserId" {
		// get attachment details from blockchain
		return getContractDetailsByUserId(stub, args)
	}

	return nil, nil
}

func main() {
	err := shim.Start(new(DTC_Chaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
