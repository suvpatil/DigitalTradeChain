package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

func InitializeChaincode(stub shim.ChaincodeStubInterface, args []string) error {
	return CreateDatabase(stub, args)
}

func SaveDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
}

func GetContract(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

}

func UpdateContractStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

}
