package main

import (
 "github.com/hyperledger/fabric/core/chaincode/shim"
 "errors"
 "encoding/json"
 "math/rand"
 "strconv"
)	

func InitializeChaincode(stub shim.ChaincodeStubInterface, args []string) error {
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

	ok, err = insertContractDetails(stub, contractsId, contractDetails)
	if !ok && err == nil {
		return nil, errors.New("Error in adding OrderDetails record")
	}

	return nil, err
}

func GetContract(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

}

func UpdateContractStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

}
